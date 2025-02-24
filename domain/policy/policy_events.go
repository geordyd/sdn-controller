package policy

type RuleAdded struct {
	PolicyID string
	Rule     Rule
}

type RuleRemoved struct {
	PolicyID string
	Rule     Rule
}
