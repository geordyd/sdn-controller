package services

import (
	"sdn/domain/policy"
	"sdn/domain/traffic"
)

func CheckPolicy(traffic traffic.Traffic) bool {
	if policy.PolicyStore.Rules[traffic.DestinationPort] {
		return true
	} else {
		// Assuming the default rule is to deny traffic
		return false
	}
}
