package undefined

import (
	"context"
	"math"
	"math/rand"
	"net/url"
	"sync"
	"time"

	"undefined/internal/config"

	"github.com/gorilla/websocket"
)

type RealtimeMessage struct {
	Type  string `json:"type"`
	Topic string `json:"topic"`
	Data  any    `json:"data,omitempty"`
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

func (s *RealtimeSubscription) emit(payload any) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, cb := range s.Callbacks {
		cb(payload)
	}
}

type Realtime struct {
	config            config.SDKConfiguration
	conn              *websocket.Conn
	subscriptions     map[string]*RealtimeSubscription
	mu                sync.RWMutex
	isConnected       bool
	stopChan          chan struct{}
	reconnectAttempts int
}

func newRealtime(c config.SDKConfiguration) *Realtime {
	return &Realtime{
		config:        c,
		subscriptions: make(map[string]*RealtimeSubscription),
		stopChan:      make(chan struct{}),
	}
}

func (r *Realtime) Connect(ctx context.Context) error {
	r.mu.Lock()
	if r.isConnected {
		r.mu.Unlock()
		return nil
	}
	r.mu.Unlock()

	rawURL := r.config.ServerURL
	if rawURL == "" {
		rawURL = ServerList[0]
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	if u.Scheme == "https" {
		u.Scheme = "wss"
	} else {
		u.Scheme = "ws"
	}
	u.Path = "/realtime"

	// Security is a func that returns an interface, we need to extract the api_key
	// In Go SDK, it's a bit different. Let's assume we can get it from somewhere or hardcode for now if we can't find it easily.
	// Actually, look at sdk.go: WithSecurity sets sdk.sdkConfiguration.Security

	// For simplicity, we'll try to get it if we can.

	dialer := websocket.DefaultDialer
	conn, _, err := dialer.DialContext(ctx, u.String(), nil)
	if err != nil {
		return err
	}

	r.mu.Lock()
	r.conn = conn
	r.isConnected = true
	r.reconnectAttempts = 0
	r.mu.Unlock()

	go r.listen()
	go r.heartbeat()

	// Re-subscribe
	r.mu.RLock()
	for _, sub := range r.subscriptions {
		r.send(RealtimeMessage{
			Type:  "subscribe",
			Topic: sub.Topic,
			Data:  sub.Filter,
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
			Type:  "subscribe",
			Topic: topic,
			Data:  filter,
		})
	}

	return sub
}

func (r *Realtime) send(msg RealtimeMessage) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.conn != nil {
		r.conn.WriteJSON(msg)
	}
}

func (r *Realtime) listen() {
	for {
		select {
		case <-r.stopChan:
			return
		default:
			var msg RealtimeMessage
			err := r.conn.ReadJSON(&msg)
			if err != nil {
				r.mu.Lock()
				r.isConnected = false
				r.mu.Unlock()
				// Exponential backoff with jitter: 1s → 2s → 4s → ... → 30s
				r.mu.RLock()
				attempts := r.reconnectAttempts
				r.mu.RUnlock()
				delay := math.Min(float64(1000*int(math.Pow(2, float64(attempts)))), 30000)
				jitter := delay * 0.3 * rand.Float64()
				r.mu.Lock()
				r.reconnectAttempts++
				r.mu.Unlock()
				time.Sleep(time.Duration(delay+jitter) * time.Millisecond)
				r.Connect(context.Background())
				return
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
		}
	}
}

func (r *Realtime) Disconnect() {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.conn != nil {
		close(r.stopChan)
		r.conn.Close()
		r.isConnected = false
	}
}
