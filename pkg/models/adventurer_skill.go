package models

type AdventurerSkill struct {
	SkillRecord
	Level       int
	Advancement struct {
		Passes   int
		Failures int
	}
}
