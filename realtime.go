package aerostack

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// RealtimeMessage matches API expected format: subscribe uses filter, chat uses roomId+text at top level.
type RealtimeMessage struct {
	Type   string         `json:"type"`
	Topic  string         `json:"topic,omitempty"`
	Filter map[string]any `json:"filter,omitempty"`
	Data   any            `json:"data,omitempty"`
}

// RealtimePayload is the typed payload dispatched to subscription callbacks.
// It carries both DB-change fields (Operation, Old) and custom-event fields (Event).
type RealtimePayload struct {
	Type      string `json:"type"`
	Topic     string `json:"topic,omitempty"`
	Operation string `json:"operation,omitempty"`
	Event     string `json:"event,omitempty"`
	Data      any    `json:"data,omitempty"`
	Old       any    `json:"old,omitempty"`
	UserID    string `json:"userId,omitempty"`
	Timestamp any    `json:"timestamp,omitempty"`
}

// HistoryMessage represents a persisted message returned from the REST history endpoint.
type HistoryMessage struct {
	ID        string `json:"id"`
	RoomID    string `json:"room_id"`
	UserID    string `json:"user_id"`
	Event     string `json:"event"`
	Data      any    `json:"data"`
	CreatedAt int64  `json:"created_at"`
}

type RealtimeSubscription struct {
	Topic     string
	Filter    map[string]any
	Callbacks []func(RealtimePayload)
	// legacyCallbacks keeps backward-compatible untyped callbacks registered via On.
	legacyCallbacks []func(any)
	mu              sync.RWMutex
	realtime        *Realtime
}

// On registers a legacy (untyped) callback that receives the raw data payload.
// For typed event handling, use OnEvent instead.
func (s *RealtimeSubscription) On(callback func(any)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.legacyCallbacks = append(s.legacyCallbacks, callback)
}

// OnEvent registers a typed callback that receives the full RealtimePayload,
// including Operation and Event fields for fine-grained routing.
func (s *RealtimeSubscription) OnEvent(callback func(RealtimePayload)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Callbacks = append(s.Callbacks, callback)
}

func (s *RealtimeSubscription) Unsubscribe(r *Realtime) {
	r.send(RealtimeMessage{Type: "unsubscribe", Topic: s.Topic})
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Callbacks = nil
	s.legacyCallbacks = nil
	// Remove from parent map to prevent memory leaks and stale message delivery
	r.removeSubscription(s.Topic)
}

func (s *RealtimeSubscription) emit(payload RealtimePayload) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, cb := range s.Callbacks {
		cb(payload)
	}
	// Fire legacy callbacks with just the data field for backward compatibility.
	for _, cb := range s.legacyCallbacks {
		cb(payload.Data)
	}
}

// Publish sends a custom event to all subscribers on this channel.
// If persist is true the server will store the message for later retrieval via GetHistory.
func (s *RealtimeSubscription) Publish(event string, data interface{}, persist bool) {
	s.mu.RLock()
	rt := s.realtime
	s.mu.RUnlock()
	if rt == nil {
		return
	}
	rt.sendMap(map[string]any{
		"type":    "publish",
		"topic":   s.Topic,
		"event":   event,
		"data":    data,
		"persist": persist,
		"id":      rt.generateID(),
	})
}

// Track broadcasts the caller's presence state to all subscribers on this channel.
// The state map is automatically synced to other clients listening for presence events.
func (s *RealtimeSubscription) Track(state map[string]interface{}) {
	s.mu.RLock()
	rt := s.realtime
	s.mu.RUnlock()
	if rt == nil {
		return
	}
	rt.sendMap(map[string]any{
		"type":  "track",
		"topic": s.Topic,
		"state": state,
	})
}

// Untrack stops broadcasting presence for this channel.
func (s *RealtimeSubscription) Untrack() {
	s.mu.RLock()
	rt := s.realtime
	s.mu.RUnlock()
	if rt == nil {
		return
	}
	rt.sendMap(map[string]any{
		"type":  "untrack",
		"topic": s.Topic,
	})
}

// GetHistory fetches persisted message history for this channel via the REST API.
// Only messages published with persist=true are returned.
// limit controls the maximum number of messages (capped server-side).
// before, if > 0, returns only messages created before this unix-ms timestamp (for pagination).
func (s *RealtimeSubscription) GetHistory(limit int, before int64) ([]HistoryMessage, error) {
	s.mu.RLock()
	rt := s.realtime
	s.mu.RUnlock()
	if rt == nil {
		return nil, fmt.Errorf("subscription is not attached to a realtime instance")
	}
	return rt.fetchHistory(s.Topic, limit, before)
}

type RealtimeStatus string

const (
	StatusIdle         RealtimeStatus = "idle"
	StatusConnecting   RealtimeStatus = "connecting"
	StatusConnected    RealtimeStatus = "connected"
	StatusReconnecting RealtimeStatus = "reconnecting"
	StatusDisconnected RealtimeStatus = "disconnected"
)

type Realtime struct {
	cfg               *Configuration
	conn              *websocket.Conn
	subscriptions     map[string]*RealtimeSubscription
	mu                sync.RWMutex
	isConnected       bool
	stopChan          chan struct{}
	stopOnce          sync.Once
	reconnectAttempts int
	// B2: Status tracking
	status          RealtimeStatus
	statusListeners []func(RealtimeStatus)
	// B3: Pong tracking
	lastPong time.Time
	// A16: Max reconnect
	maxReconnectAttempts int
	maxRetriesListeners  []func()
	// Context saved from Connect for API key access in REST calls
	ctx context.Context
}

func newRealtime(cfg *Configuration) *Realtime {
	return &Realtime{
		cfg:                  cfg,
		subscriptions:        make(map[string]*RealtimeSubscription),
		stopChan:             make(chan struct{}),
		status:               StatusIdle,
		maxReconnectAttempts: -1, // -1 = infinite
	}
}

// SetMaxReconnectAttempts sets the max number of reconnect attempts (0 = infinite).
func (r *Realtime) SetMaxReconnectAttempts(n int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.maxReconnectAttempts = n
}

func (r *Realtime) Status() RealtimeStatus {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.status
}

func (r *Realtime) OnStatusChange(cb func(RealtimeStatus)) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.statusListeners = append(r.statusListeners, cb)
}

func (r *Realtime) OnMaxRetriesExceeded(cb func()) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.maxRetriesListeners = append(r.maxRetriesListeners, cb)
}

func (r *Realtime) setStatus(s RealtimeStatus) {
	r.status = s
	for _, cb := range r.statusListeners {
		cb(s)
	}
}

// SetToken updates the auth token on a live connection (B4). API expects token at top level.
func (r *Realtime) SetToken(newToken string) {
	r.sendMap(map[string]any{"type": "auth", "token": newToken})
}

// generateID produces a simple unique message ID (random hex + timestamp) for ack tracking.
func (r *Realtime) generateID() string {
	return fmt.Sprintf("%x%x", rand.Int63(), time.Now().UnixNano())
}

func (r *Realtime) Connect(ctx context.Context) error {
	r.mu.Lock()
	if r.isConnected {
		r.mu.Unlock()
		return nil
	}
	r.setStatus(StatusConnecting)
	r.ctx = ctx
	r.mu.Unlock()

	rawURL, err := r.cfg.ServerURLWithContext(ctx, "")
	if err != nil {
		return err
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return err
	}
	u.Path = ""
	if u.Scheme == "https" {
		u.Scheme = "wss"
	} else if u.Scheme == "http" {
		u.Scheme = "ws"
	}
	u.Path = "/api/realtime"

	// SECURITY: Pass API key via Sec-WebSocket-Protocol header — never as URL query param
	// (URL params appear in CDN logs, browser history, and Referer headers).
	var dialHeaders http.Header
	if ctx != nil {
		if apiKeys := ctx.Value(ContextAPIKeys); apiKeys != nil {
			if keys, ok := apiKeys.(map[string]APIKey); ok {
				if key, ok := keys["ApiKeyAuth"]; ok {
					dialHeaders = http.Header{}
					dialHeaders.Set("Sec-WebSocket-Protocol", fmt.Sprintf("aerostack-key.%s, aerostack-v1", key.Key))
				}
			}
		}
	}

	dialer := websocket.DefaultDialer
	conn, _, err := dialer.DialContext(ctx, u.String(), dialHeaders)
	if err != nil {
		r.mu.Lock()
		r.setStatus(StatusDisconnected)
		r.mu.Unlock()
		return err
	}

	r.mu.Lock()
	r.conn = conn
	r.isConnected = true
	r.reconnectAttempts = 0
	r.lastPong = time.Now()
	r.stopChan = make(chan struct{})
	r.stopOnce = sync.Once{}
	r.setStatus(StatusConnected)
	r.mu.Unlock()

	go r.listen(ctx)
	go r.heartbeat()

	// Re-subscribe (API expects filter, not data)
	r.mu.RLock()
	for _, sub := range r.subscriptions {
		r.send(RealtimeMessage{
			Type:   "subscribe",
			Topic:  sub.Topic,
			Filter: sub.Filter,
		})
	}
	r.mu.RUnlock()

	return nil
}

func (r *Realtime) Channel(topic string, filter map[string]any) *RealtimeSubscription {
	r.mu.Lock()
	defer r.mu.Unlock()

	if sub, ok := r.subscriptions[topic]; ok {
		return sub
	}

	sub := &RealtimeSubscription{
		Topic:    topic,
		Filter:   filter,
		realtime: r,
	}
	r.subscriptions[topic] = sub

	if r.isConnected {
		r.send(RealtimeMessage{
			Type:   "subscribe",
			Topic:  topic,
			Filter: filter,
		})
	}

	return sub
}

// SendChat sends a chat message to a room. API expects roomId and text at top level.
func (r *Realtime) SendChat(roomID, text string) {
	r.sendMap(map[string]any{"type": "chat", "roomId": roomID, "text": text})
}

// removeSubscription deletes a subscription from the internal map (called on Unsubscribe).
func (r *Realtime) removeSubscription(topic string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.subscriptions, topic)
}

func (r *Realtime) send(msg RealtimeMessage) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.conn != nil {
		_ = r.conn.WriteJSON(msg)
	}
}

func (r *Realtime) sendMap(m map[string]any) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.conn != nil {
		data, _ := json.Marshal(m)
		_ = r.conn.WriteMessage(websocket.TextMessage, data)
	}
}

// apiKey extracts the API key string from the saved context, if available.
func (r *Realtime) apiKey() string {
	r.mu.RLock()
	ctx := r.ctx
	r.mu.RUnlock()
	if ctx == nil {
		return ""
	}
	if apiKeys := ctx.Value(ContextAPIKeys); apiKeys != nil {
		if keys, ok := apiKeys.(map[string]APIKey); ok {
			if key, ok := keys["ApiKeyAuth"]; ok {
				return key.Key
			}
		}
	}
	return ""
}

// httpBaseURL derives the HTTP base URL from the Configuration (strips /v1 suffix).
func (r *Realtime) httpBaseURL() (string, error) {
	r.mu.RLock()
	ctx := r.ctx
	r.mu.RUnlock()
	rawURL, err := r.cfg.ServerURLWithContext(ctx, "")
	if err != nil {
		return "", err
	}
	// Strip trailing /v1 to get the base URL used by REST endpoints.
	base := strings.TrimSuffix(rawURL, "/v1")
	base = strings.TrimSuffix(base, "/v1/")
	return base, nil
}

// fetchHistory calls the REST API to retrieve persisted message history for a room/topic.
func (r *Realtime) fetchHistory(room string, limit int, before int64) ([]HistoryMessage, error) {
	base, err := r.httpBaseURL()
	if err != nil {
		return nil, fmt.Errorf("failed to resolve server URL: %w", err)
	}

	u, err := url.Parse(base + "/api/v1/public/realtime/history")
	if err != nil {
		return nil, fmt.Errorf("failed to build history URL: %w", err)
	}
	q := u.Query()
	q.Set("room", room)
	if limit > 0 {
		q.Set("limit", fmt.Sprintf("%d", limit))
	}
	if before > 0 {
		q.Set("before", fmt.Sprintf("%d", before))
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create history request: %w", err)
	}

	apiKey := r.apiKey()
	if apiKey != "" {
		req.Header.Set("X-Aerostack-Key", apiKey)
	}

	client := r.cfg.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("history request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read history response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("history endpoint returned status %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Messages []HistoryMessage `json:"messages"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse history response: %w", err)
	}
	return result.Messages, nil
}

func (r *Realtime) listen(ctx context.Context) {
	for {
		select {
		case <-r.stopChan:
			return
		default:
			_, rawBytes, err := r.conn.ReadMessage()
			if err != nil {
				select {
				case <-r.stopChan:
					return
				default:
				}

				r.mu.Lock()
				r.isConnected = false
				attempts := r.reconnectAttempts
				r.reconnectAttempts++
				r.setStatus(StatusReconnecting)
				r.mu.Unlock()

				// A16: Max retries check
				r.mu.RLock()
				maxAttempts := r.maxReconnectAttempts
				r.mu.RUnlock()
				if maxAttempts >= 0 && attempts >= maxAttempts {
					r.mu.Lock()
					r.setStatus(StatusDisconnected)
					listeners := make([]func(), len(r.maxRetriesListeners))
					copy(listeners, r.maxRetriesListeners)
					r.mu.Unlock()
					for _, cb := range listeners {
						cb()
					}
					return
				}

				delay := math.Min(float64(1000*int(math.Pow(2, float64(attempts)))), 30000)
				jitter := delay * 0.3 * rand.Float64()

				select {
				case <-ctx.Done():
					return
				case <-time.After(time.Duration(delay+jitter) * time.Millisecond):
					r.Connect(ctx)
				}
				return
			}

			// Parse the raw message into a generic map so we can inspect the type
			// and extract all fields (including custom ones like "event", "operation", etc.).
			var raw map[string]any
			if err := json.Unmarshal(rawBytes, &raw); err != nil {
				continue
			}

			msgType, _ := raw["type"].(string)
			msgTopic, _ := raw["topic"].(string)

			// B3: Track pong
			if msgType == "pong" {
				r.mu.Lock()
				r.lastPong = time.Now()
				r.mu.Unlock()
				continue
			}

			// Handle ack messages (server acknowledgement for published messages).
			if msgType == "ack" {
				continue
			}

			// Server confirmed subscription with its normalized topic (e.g. 'table/orders/<projectId>').
			// Re-key our subscription map so incoming db_change messages can be routed correctly.
			if msgType == "subscribed" && msgTopic != "" {
				r.mu.Lock()
				for origTopic, sub := range r.subscriptions {
					if msgTopic != origTopic && strings.HasPrefix(msgTopic, origTopic) {
						delete(r.subscriptions, origTopic)
						sub.Topic = msgTopic
						r.subscriptions[msgTopic] = sub
						break
					}
				}
				r.mu.Unlock()
				continue
			}

			// Route db_change, chat_message, and custom event messages to subscriptions.
			if msgType == "db_change" || msgType == "chat_message" || msgType == "event" {
				payload := RealtimePayload{
					Type:  msgType,
					Topic: msgTopic,
				}
				if op, ok := raw["operation"].(string); ok {
					payload.Operation = op
				}
				if ev, ok := raw["event"].(string); ok {
					payload.Event = ev
				}
				if d, ok := raw["data"]; ok {
					payload.Data = d
				}
				if old, ok := raw["old"]; ok {
					payload.Old = old
				}
				if uid, ok := raw["userId"].(string); ok {
					payload.UserID = uid
				}
				if ts, ok := raw["timestamp"]; ok {
					payload.Timestamp = ts
				}

				r.mu.RLock()
				if sub, ok := r.subscriptions[msgTopic]; ok {
					sub.emit(payload)
				}
				r.mu.RUnlock()
				continue
			}

			// Fallback: route any other message with a topic to its subscription.
			if msgTopic != "" {
				payload := RealtimePayload{
					Type:  msgType,
					Topic: msgTopic,
					Data:  raw["data"],
				}
				r.mu.RLock()
				if sub, ok := r.subscriptions[msgTopic]; ok {
					sub.emit(payload)
				}
				r.mu.RUnlock()
			}
		}
	}
}

func (r *Realtime) heartbeat() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-r.stopChan:
			return
		case <-ticker.C:
			r.send(RealtimeMessage{Type: "ping"})
			// B3: Dead connection check — no pong in 70s
			r.mu.RLock()
			elapsed := time.Since(r.lastPong)
			r.mu.RUnlock()
			if !r.lastPong.IsZero() && elapsed > 70*time.Second {
				r.mu.Lock()
				if r.conn != nil {
					r.conn.Close()
				}
				r.mu.Unlock()
			}
		}
	}
}

// Disconnect closes the WebSocket and stops all goroutines. Safe to call multiple times.
func (r *Realtime) Disconnect() {
	r.stopOnce.Do(func() {
		close(r.stopChan)
	})
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.conn != nil {
		r.conn.Close()
		r.conn = nil
		r.isConnected = false
	}
	r.setStatus(StatusDisconnected)
}
