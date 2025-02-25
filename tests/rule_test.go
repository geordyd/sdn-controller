package main

import (
	"sdn/domain/policy"
	"testing"

	"github.com/google/uuid"
)

func TestAddRule(t *testing.T) {
	p := policy.NewPolicy(uuid.New().String())

	rule := policy.Rule{
		Port:   80,
		Action: "allow",
	}

	p.AddRule(rule)

	if len(p.Rules) != 1 {
		t.Errorf("expected 1 rule, got %v", len(p.Rules))
	}

	if p.Rules[rule.Port] != rule.Action {
		t.Errorf("rule was not added")
	}
}

func TestRemoveRule(t *testing.T) {
	p := policy.NewPolicy(uuid.New().String())

	rule := policy.Rule{
		Port:   80,
		Action: "allow",
	}

	p.Rules[rule.Port] = rule.Action

	p.RemoveRule(rule)

	if p.Rules[rule.Port] != "" {
		t.Errorf("rule was not removed")
	}
}
