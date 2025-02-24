package handlers

import (
	"fmt"
	"sdn/application/services"
	"sdn/domain/traffic"
	"sdn/infrastructure/pubsub"
)

func TrafficReceivedHandler(ch <-chan pubsub.Event) {

	eventPublisher := pubsub.NewEventPublisher(pubsub.Instance)

	for event := range ch {
		trafficData, ok := event.Data.(traffic.TrafficReceived)
		if !ok {
			fmt.Println("Invalid event data")
			continue
		}
		fmt.Printf("%s: Type: %s, Source IP: %s, Destination Port: %d\n",
			event.Timestamp.Format("2006-01-02 15:04:05"),
			event.Type,
			trafficData.Packet.SourceIP,
			trafficData.Packet.DestinationPort)

		if services.CheckPolicy(trafficData.Packet) {
			trafficAllowedEvent := traffic.TrafficAllowed(trafficData)
			eventPublisher.PublishEvent("TrafficAllowed", trafficAllowedEvent)
		} else {
			trafficBlockedEvent := traffic.TrafficBlocked(trafficData)
			eventPublisher.PublishEvent("TrafficBlocked", trafficBlockedEvent)
		}
	}
}

func TrafficAllowedHandler(ch <-chan pubsub.Event) {
	for event := range ch {
		traffic, ok := event.Data.(traffic.TrafficAllowed)
		if !ok {
			fmt.Println("Invalid event data")
			continue
		}
		fmt.Printf("%s: Type: %s, Source IP: %s, Destination Port: %d\n",
			event.Timestamp.Format("2006-01-02 15:04:05"),
			event.Type,
			traffic.Packet.SourceIP,
			traffic.Packet.DestinationPort)
	}
}

func TrafficBlockedHandler(ch <-chan pubsub.Event) {
	for event := range ch {
		traffic, ok := event.Data.(traffic.TrafficBlocked)
		if !ok {
			fmt.Println("Invalid event data")
			continue
		}
		fmt.Printf("%s: Type: %s, Source IP: %s, Destination Port: %d\n",
			event.Timestamp.Format("2006-01-02 15:04:05"),
			event.Type,
			traffic.Packet.SourceIP,
			traffic.Packet.DestinationPort)
	}
}
