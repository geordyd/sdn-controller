package pubsub

import "time"

var EventStore []Event

type Event struct {
	Type      string
	Timestamp time.Time
	Data      any
}

type PubSub struct {
	subscribers map[string][]chan<- Event
}

func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]chan<- Event),
	}
}

func (ps *PubSub) Subscribe(eventType string) <-chan Event {
	ch := make(chan Event)
	ps.subscribers[eventType] = append(ps.subscribers[eventType], ch)
	return ch
}

func (ps *PubSub) Publish(event Event) {
	EventStore = append(EventStore, event)
	for _, ch := range ps.subscribers[event.Type] {
		ch <- event
	}
}

type EventPublisher struct {
	PubSub *PubSub
}

func NewEventPublisher(pubSub *PubSub) *EventPublisher {
	return &EventPublisher{
		PubSub: pubSub,
	}
}

func (ep *EventPublisher) PublishEvent(eventType string, data any) {
	event := Event{
		Type:      eventType,
		Timestamp: time.Now(),
		Data:      data,
	}

	ep.PubSub.Publish(event)
}

var Instance *PubSub
