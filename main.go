package main

import (
	"sdn/domain/traffic"
	"sdn/handlers"
	"sdn/pubsub"
	"time"
)

func main() {
	pubsub.Instance = pubsub.NewPubSub()

	message := pubsub.Instance.Subscribe("TrafficReceived")

	go handlers.TrafficReceivedHandler(message)

	eventPublisher := pubsub.NewEventPublisher(pubsub.Instance)

	go func() {
		for {
			trafficEvent := traffic.TrafficReceived{
				Packet: traffic.Traffic{
					SourceIP:        "",
					DestinationPort: 0,
				},
			}
			eventPublisher.PublishEvent("TrafficReceived", trafficEvent)

			time.Sleep(2 * time.Second)
		}
	}()

	select {}
}
