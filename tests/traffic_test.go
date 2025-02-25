package main

import (
	"sdn/application/handlers"
	"sdn/domain/policy"
	"sdn/domain/traffic"
	"sdn/infrastructure/pubsub"
	"testing"
	"time"

	"github.com/google/uuid"
)

func setup() (*pubsub.PubSub, *pubsub.EventPublisher, *policy.Policy) {
	policyStore := policy.NewPolicy(uuid.New().String())
	ps := pubsub.NewPubSub()
	eventPublisher := pubsub.NewEventPublisher(ps)
	return ps, eventPublisher, policyStore
}

func TestTrafficAllowed(t *testing.T) {
	ps, eventPublisher, policyStore := setup()

	rule := policy.Rule{
		Port:   80,
		Action: "allow",
	}

	policyStore.Rules[rule.Port] = rule.Action

	trafficReceived := ps.Subscribe("TrafficReceived")
	trafficAllowed := ps.Subscribe("TrafficAllowed")

	go handlers.TrafficReceivedHandler(trafficReceived, eventPublisher, policyStore)
	go handlers.TrafficAllowedHandler(trafficAllowed)

	eventPublisher.PublishEvent("TrafficReceived", traffic.TrafficReceived{
		Packet: traffic.Traffic{
			SourceIP:        "10.13.37.1",
			DestinationPort: 80,
		},
	})

	// Wait for the event to be processed
	time.Sleep(1 * time.Second)

	if len(ps.EventStore) != 2 {
		t.Errorf("expected 2 events, got %v", len(ps.EventStore))
	}

	if eventType := ps.EventStore[len(ps.EventStore)-1].Type; eventType != "TrafficAllowed" {
		t.Errorf("expected TrafficAllowed, got %s", eventType)
	}
}

func TestTrafficBlocked(t *testing.T) {

	ps, eventPublisher, policyStore := setup()

	rule := policy.Rule{
		Port:   80,
		Action: "deny",
	}

	policyStore.Rules[rule.Port] = rule.Action

	ps = pubsub.NewPubSub()
	eventPublisher = pubsub.NewEventPublisher(ps)

	trafficReceived := ps.Subscribe("TrafficReceived")
	trafficBlocked := ps.Subscribe("TrafficBlocked")

	go handlers.TrafficReceivedHandler(trafficReceived, eventPublisher, policyStore)
	go handlers.TrafficBlockedHandler(trafficBlocked)

	eventPublisher.PublishEvent("TrafficReceived", traffic.TrafficReceived{
		Packet: traffic.Traffic{
			SourceIP:        "10.13.37.1",
			DestinationPort: 80,
		},
	})

	// Wait for the event to be processed
	time.Sleep(1 * time.Second)

	if len(ps.EventStore) != 2 {
		t.Errorf("expected 2 events, got %v", len(ps.EventStore))
	}

	if eventType := ps.EventStore[len(ps.EventStore)-1].Type; eventType != "TrafficBlocked" {
		t.Errorf("expected TrafficBlocked, got %s", eventType)
	}
}

func TestTrafficDropped(t *testing.T) {

	ps, eventPublisher, policyStore := setup()

	rule := policy.Rule{
		Port:   80,
		Action: "drop",
	}

	policyStore.Rules[rule.Port] = rule.Action

	trafficReceived := ps.Subscribe("TrafficReceived")
	trafficDropped := ps.Subscribe("TrafficDropped")

	go handlers.TrafficReceivedHandler(trafficReceived, eventPublisher, policyStore)
	go handlers.TrafficDroppedHandler(trafficDropped)

	eventPublisher.PublishEvent("TrafficReceived", traffic.TrafficReceived{
		Packet: traffic.Traffic{
			SourceIP:        "10.13.37.1",
			DestinationPort: 80,
		},
	})

	// Wait for the event to be processed
	time.Sleep(1 * time.Second)

	if len(ps.EventStore) != 2 {
		t.Errorf("expected 2 events, got %v", len(ps.EventStore))
	}

	if eventType := ps.EventStore[len(ps.EventStore)-1].Type; eventType != "TrafficDropped" {
		t.Errorf("expected TrafficDropped, got %s", eventType)
	}
}
