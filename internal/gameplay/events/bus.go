package events

import "reflect"

type Handler func(Event)

type Bus struct {
	handlers map[reflect.Type][]Handler
	queue    []Event
}

func NewBus() *Bus {
	return &Bus{
		handlers: make(map[reflect.Type][]Handler),
		queue:    make([]Event, 0),
	}
}

// Subscribe adds a handler for a specific event type
func (b *Bus) Subscribe(eventType reflect.Type, handler Handler) {
	b.handlers[eventType] = append(b.handlers[eventType], handler)
}

// Publish adds an event to the queue
func (b *Bus) Publish(e Event) {
	b.queue = append(b.queue, e)
}

// Process dispatches all queued events to their subscribers
func (b *Bus) Process() {
	for len(b.queue) > 0 {
		currentQueue := b.queue
		b.queue = nil
		for _, event := range currentQueue {
			if handlers, ok := b.handlers[reflect.TypeOf(event)]; ok {
				for _, handler := range handlers {
					handler(event)
				}
			}
		}
	}
}
