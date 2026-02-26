package aerostack

import (
	"context"
	"encoding/json"
	"math"
	"math/rand"
	"net/url"
	"sync"
	"time"

	"github.com/aerostackdev/sdks/packages/go/internal/config"
	"github.com/aerostackdev/sdks/packages/go/pkg/models/shared"
	"github.com/gorilla/websocket"
)

// RealtimeMessage matches API expected format: subscribe uses filter, chat uses roomId+text at top level.
type RealtimeMessage struct {
	Type   string         `json:"type"`
	Topic  string         `json:"topic,omitempty"`
	Filter map[string]any `json:"filter,omitempty"`
	Data   any            `json:"data,omitempty"`
}

type RealtimeSubscription struct {
	Topic     string
	Filter    map[string]any
	Callbacks []func(any)
	mu        sync.RWMutex
}

func (s *RealtimeSubscription) On(callback func(any)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Callbacks = append(s.Callbacks, callback)
}

func (s *RealtimeSubscription) Unsubscribe(r *Realtime) {
	r.send(RealtimeMessage{Type: "unsubscribe", Topic: s.Topic})
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Callbacks = nil
}

func (s *RealtimeSubscription) emit(payload any) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, cb := range s.Callbacks {
		cb(payload)
	}
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
	config            config.SDKConfiguration
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
}

func newRealtime(c config.SDKConfiguration) *Realtime {
	return &Realtime{
		config:               c,
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

func (r *Realtime) Connect(ctx context.Context) error {
	r.mu.Lock()
	if r.isConnected {
		r.mu.Unlock()
		return nil
	}
	r.setStatus(StatusConnecting)
	r.mu.Unlock()

	rawURL := r.config.ServerURL
	if rawURL == "" {
		rawURL = ServerList[0]
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return err
	}
	u.Path = ""
	if u.Scheme == "https" {
		u.Scheme = "wss"
	} else {
		u.Scheme = "ws"
	}
	u.Path = "/api/realtime"

	// Attach auth params
	if r.config.Security != nil {
		secIface, err := r.config.Security(ctx)
		if err == nil && secIface != nil {
			if sec, ok := secIface.(shared.Security); ok {
				q := u.Query()
				if sec.APIKeyAuth != "" {
					q.Set("apiKey", sec.APIKeyAuth)
				}
				u.RawQuery = q.Encode()
			}
		}
	}

	dialer := websocket.DefaultDialer
	conn, _, err := dialer.DialContext(ctx, u.String(), nil)
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
		Topic:  topic,
		Filter: filter,
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

func (r *Realtime) listen(ctx context.Context) {
	for {
		select {
		case <-r.stopChan:
			return
		default:
			var msg RealtimeMessage
			err := r.conn.ReadJSON(&msg)
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

			// B3: Track pong
			if msg.Type == "pong" {
				r.mu.Lock()
				r.lastPong = time.Now()
				r.mu.Unlock()
				continue
			}

			r.mu.RLock()
			if sub, ok := r.subscriptions[msg.Topic]; ok {
				sub.emit(msg.Data)
			}
			r.mu.RUnlock()
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
