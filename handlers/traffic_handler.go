package handlers

import (
	"fmt"
	"sdn/domain/traffic"
	"sdn/pubsub"
)

func TrafficReceivedHandler(ch <-chan pubsub.Event) {
	for event := range ch {
		traffic, ok := event.Data.(traffic.TrafficReceived)
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
