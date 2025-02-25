package handlers

import (
	"fmt"
	"sdn/domain/policy"
	"sdn/infrastructure/pubsub"
)

func RuleAddedHandler(ch <-chan pubsub.Event) {
	for event := range ch {
		rule, ok := event.Data.(policy.RuleAdded)
		if !ok {
			fmt.Println("Invalid event data")
			continue
		}
		fmt.Printf("%s: Type: %s, Allowed: %v, Port: %d\n",
			event.Timestamp.Format("2006-01-02 15:04:05"),
			event.Type,
			rule.Rule.Action,
			rule.Rule.Port)
	}
}

func RuleRemovedHandler(ch <-chan pubsub.Event) {
	for event := range ch {
		rule, ok := event.Data.(policy.RuleRemoved)
		if !ok {
			fmt.Println("Invalid event data")
			continue
		}
		fmt.Printf("%s: Type: %s, Allowed: %v, Port: %d\n",
			event.Timestamp.Format("2006-01-02 15:04:05"),
			event.Type,
			rule.Rule.Action,
			rule.Rule.Port)
	}
}
