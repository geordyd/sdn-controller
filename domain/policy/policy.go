package policy

type Rule struct {
	Port   int
	Action string
}

type Policy struct {
	ID    string
	Rules map[int]string
}

func NewPolicy(id string) *Policy {
	return &Policy{
		ID:    id,
		Rules: make(map[int]string),
	}
}

func (p *Policy) AddRule(rule Rule) {
	p.Rules[rule.Port] = rule.Action
}

func (p *Policy) RemoveRule(rule Rule) {
	delete(p.Rules, rule.Port)
}
