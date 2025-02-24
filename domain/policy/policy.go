package policy

type Rule struct {
	Port    int
	Allowed bool
}

type Policy struct {
	ID    string
	Rules map[int]bool
}

var PolicyStore Policy

func NewPolicy(id string) {
	PolicyStore = Policy{
		ID:    id,
		Rules: make(map[int]bool),
	}
}

func (p *Policy) AddRule(rule Rule) {
	p.Rules[rule.Port] = rule.Allowed
}

func (p *Policy) RemoveRule(rule Rule) {
	delete(p.Rules, rule.Port)
}
