package pubsub

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID        uuid.UUID
	Type      string
	Timestamp time.Time
	Data      any
}

type PubSub struct {
	subscribers map[string][]chan<- Event
	EventStore  []*Event
}

func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]chan<- Event),
		EventStore:  []*Event{},
	}
}

func (ps *PubSub) Subscribe(eventType string) <-chan Event {
	ch := make(chan Event)
	ps.subscribers[eventType] = append(ps.subscribers[eventType], ch)
	return ch
}

func (ps *PubSub) Publish(event *Event) {
	ps.EventStore = append(ps.EventStore, event)
	for _, ch := range ps.subscribers[event.Type] {
		ch <- *event
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
		ID:        uuid.New(),
		Type:      eventType,
		Timestamp: time.Now(),
		Data:      data,
	}

	ep.PubSub.Publish(&event)
}
