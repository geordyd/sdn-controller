package policy

// value object
type Rule struct {
	Port    int
	Allowed bool
}

// aggregate
type Policy struct {
	ID    string
	Rules map[int]bool
}

var PolicyStore Policy

func (p *Policy) AddRule(rule Rule) {
	p.Rules[rule.Port] = rule.Allowed
}

func (p *Policy) RemoveRule(rule Rule) {
	delete(p.Rules, rule.Port)
}
