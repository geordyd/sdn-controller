package main

import (
	"fmt"
	"net/http"
	"sdn/domain/policy"
	"sdn/domain/traffic"
	"sdn/handlers"
	"sdn/pubsub"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func main() {

	policy.NewPolicy(uuid.New().String())

	pubsub.Instance = pubsub.NewPubSub()

	trafficReceived := pubsub.Instance.Subscribe("TrafficReceived")
	trafficAllowed := pubsub.Instance.Subscribe("TrafficAllowed")
	trafficBlocked := pubsub.Instance.Subscribe("TrafficBlocked")
	ruleAdded := pubsub.Instance.Subscribe("RuleAdded")
	ruleRemoved := pubsub.Instance.Subscribe("RuleRemoved")

	go handlers.TrafficReceivedHandler(trafficReceived)
	go handlers.TrafficAllowedHandler(trafficAllowed)
	go handlers.TrafficBlockedHandler(trafficBlocked)
	go handlers.RuleAddedHandler(ruleAdded)
	go handlers.RuleRemovedHandler(ruleRemoved)

	eventPublisher := pubsub.NewEventPublisher(pubsub.Instance)

	go generateTraffic("10.13.37.1", 80, 1, eventPublisher)
	go generateTraffic("10.13.37.1", 443, 5, eventPublisher)
	go generateTraffic("10.13.37.1", 23, 10, eventPublisher)

	go func() {
		mux := http.NewServeMux()

		mux.HandleFunc("/addrule/{state}/{port}", func(w http.ResponseWriter, r *http.Request) {
			var allowed bool
			if r.PathValue("state") == "allow" {
				allowed = true
			} else if r.PathValue("state") == "deny" {
				allowed = false
			} else {
				http.Error(w, "Invalid state", http.StatusBadRequest)
				return
			}

			port, err := strconv.Atoi(r.PathValue("port"))
			if err != nil {
				http.Error(w, "Invalid port", http.StatusBadRequest)
				return
			}

			rule := policy.Rule{
				Allowed: allowed,
				Port:    port,
			}

			policy.PolicyStore.AddRule(rule)

			eventPublisher.PublishEvent("RuleAdded", policy.RuleAdded{
				PolicyID: policy.PolicyStore.ID,
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

			policy.PolicyStore.RemoveRule(rule)

			eventPublisher.PublishEvent("RuleRemoved", policy.RuleRemoved{
				PolicyID: policy.PolicyStore.ID,
				Rule:     rule,
			})
		})

		mux.HandleFunc("/getevents", func(w http.ResponseWriter, r *http.Request) {
			events := pubsub.EventStore
			for _, event := range events {
				fmt.Fprintf(w, "%s: Type: %s, Data: %v\n",
					event.Timestamp.Format("2006-01-02 15:04:05"),
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
