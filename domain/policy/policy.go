package policy

type Rule struct {
	Port    int
	Allowed string
}

type Policy struct {
	ID    string
	Rules map[int]string
}

var PolicyStore Policy

func NewPolicy(id string) {
	PolicyStore = Policy{
		ID:    id,
		Rules: make(map[int]string),
	}
}

func (p *Policy) AddRule(rule Rule) {
	p.Rules[rule.Port] = rule.Allowed
}

func (p *Policy) RemoveRule(rule Rule) {
	delete(p.Rules, rule.Port)
}
