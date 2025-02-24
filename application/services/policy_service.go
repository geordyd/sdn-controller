package services

import (
	"fmt"
	"sdn/domain/policy"
	"sdn/domain/traffic"
)

func CheckPolicy(traffic traffic.Traffic) string {
	if policy.PolicyStore.Rules[traffic.DestinationPort] == "allow" {
		return "Allowed"
	} else if policy.PolicyStore.Rules[traffic.DestinationPort] == "deny" {
		fmt.Println(policy.PolicyStore.Rules[traffic.DestinationPort])
		return "Blocked"
	} else {
		// Assuming the default rule is to deny traffic
		return "Dropped"
	}
}
