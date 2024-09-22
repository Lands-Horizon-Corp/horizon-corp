package events

import (
	"sync"
)

// EventType represents the type of event
type EventType string

// Event represents an event with a type and payload
type Event struct {
	Type    EventType
	Payload interface{}
}

// Listener defines a function that handles events
type Listener func(event Event)

// Dispatcher manages event listeners and dispatching events
type Dispatcher struct {
	listeners map[EventType][]Listener
	mu        sync.RWMutex
}

var dispatcherInstance *Dispatcher
var once sync.Once

// GetDispatcher returns the singleton dispatcher instance
func GetDispatcher() *Dispatcher {
	once.Do(func() {
		dispatcherInstance = &Dispatcher{
			listeners: make(map[EventType][]Listener),
		}
	})
	return dispatcherInstance
}

// Register adds a listener for a specific event type
func (d *Dispatcher) Register(eventType EventType, listener Listener) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.listeners[eventType] = append(d.listeners[eventType], listener)
}

// Dispatch sends an event to all registered listeners for its type
func (d *Dispatcher) Dispatch(event Event) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if listeners, exists := d.listeners[event.Type]; exists {
		for _, listener := range listeners {
			// Run listeners asynchronously
			go listener(event)
		}
	}
}
