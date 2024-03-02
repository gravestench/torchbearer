package models

type SkillRecord struct {
	Name          string
	Description   string
	BeginnersLuck string
	Tools         []string
	Consumables   []string
	SupportSkills []string
	Factors       map[string][]SkillFactor
}

type SkillFactor struct {
	Value       int
	Description string
}
