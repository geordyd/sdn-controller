package services

import (
	"sdn/domain/policy"
	"sdn/domain/traffic"
)

func CheckPolicy(traffic traffic.Traffic, policyStore *policy.Policy) string {
	if policyStore.Rules[traffic.DestinationPort] == "allow" {
		return "Allowed"
	} else if policyStore.Rules[traffic.DestinationPort] == "deny" {
		return "Blocked"
	} else {
		// Assuming the default rule is to deny traffic
		return "Dropped"
	}
}
