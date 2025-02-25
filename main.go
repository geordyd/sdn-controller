package main

import (
	"fmt"
	"net/http"
	"sdn/application/handlers"
	"sdn/domain/policy"
	"sdn/domain/traffic"
	"sdn/infrastructure/pubsub"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func main() {

	policyStore := policy.NewPolicy(uuid.New().String())

	ps := pubsub.NewPubSub()
	eventPublisher := pubsub.NewEventPublisher(ps)

	trafficReceived := ps.Subscribe("TrafficReceived")
	trafficAllowed := ps.Subscribe("TrafficAllowed")
	trafficBlocked := ps.Subscribe("TrafficBlocked")
	trafficDropped := ps.Subscribe("TrafficDropped")
	ruleAdded := ps.Subscribe("RuleAdded")
	ruleRemoved := ps.Subscribe("RuleRemoved")

	go handlers.TrafficReceivedHandler(trafficReceived, eventPublisher, policyStore)
	go handlers.TrafficAllowedHandler(trafficAllowed)
	go handlers.TrafficBlockedHandler(trafficBlocked)
	go handlers.TrafficDroppedHandler(trafficDropped)
	go handlers.RuleAddedHandler(ruleAdded)
	go handlers.RuleRemovedHandler(ruleRemoved)

	go generateTraffic("10.13.37.1", 80, 1, eventPublisher)
	go generateTraffic("10.13.37.1", 443, 5, eventPublisher)
	go generateTraffic("10.13.37.1", 23, 10, eventPublisher)

	go func() {
		mux := http.NewServeMux()

		mux.HandleFunc("/addrule/{action}/{port}", func(w http.ResponseWriter, r *http.Request) {
			var allowed string
			if r.PathValue("action") == "allow" {
				allowed = r.PathValue("action")
			} else if r.PathValue("action") == "deny" {
				allowed = r.PathValue("action")
			} else {
				http.Error(w, "Invalid action", http.StatusBadRequest)
				return
			}

			port, err := strconv.Atoi(r.PathValue("port"))
			if err != nil {
				http.Error(w, "Invalid port", http.StatusBadRequest)
				return
			}

			rule := policy.Rule{
				Action: allowed,
				Port:   port,
			}

			policyStore.AddRule(rule)

			eventPublisher.PublishEvent("RuleAdded", policy.RuleAdded{
				PolicyID: policyStore.ID,
				Rule:     rule,
			})
		})

		mux.HandleFunc("/removerule/{port}", func(w http.ResponseWriter, r *http.Request) {
			port, err := strconv.Atoi(r.PathValue("port"))
			if err != nil {
				http.Error(w, "Invalid port", http.StatusBadRequest)
				return
			}

			rule := policy.Rule{
				Port: port,
			}

			policyStore.RemoveRule(rule)

			eventPublisher.PublishEvent("RuleRemoved", policy.RuleRemoved{
				PolicyID: policyStore.ID,
				Rule:     rule,
			})
		})

		mux.HandleFunc("/getevents", func(w http.ResponseWriter, r *http.Request) {
			events := ps.EventStore
			for _, event := range events {
				fmt.Fprintf(w, "%s: ID: %v, Type: %s, Data: %v\n",
					event.Timestamp.Format("2006-01-02 15:04:05"),
					event.ID,
					event.Type,
					event.Data)
			}
		})

		port := 1337
		addr := fmt.Sprintf(":%d", port)
		fmt.Printf("Server listening on http://localhost%s\n", addr)
		if err := http.ListenAndServe(addr, mux); err != nil {
			panic(err)
		}
	}()

	select {}
}

func generateTraffic(sourceIP string, destinationPort int, interval int, eventPublisher *pubsub.EventPublisher) {
	for {
		trafficEvent := traffic.TrafficReceived{
			Packet: traffic.Traffic{
				ID:              uuid.New(),
				SourceIP:        sourceIP,
				DestinationPort: destinationPort,
			},
		}
		eventPublisher.PublishEvent("TrafficReceived", trafficEvent)
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
