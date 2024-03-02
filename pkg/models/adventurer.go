package models

import (
	"github.com/google/uuid"
)

func NewAdventurer() *Adventurer {
	a := &Adventurer{}
	a.ID = uuid.New()
	a.Traits = make(map[string]*AdventurerTrait)
	a.Skills = make(map[string]*AdventurerSkill)
	a.Wises = make(map[string]*AdventurerWise)
	a.Raiment = make([]string, 0)
	return a
}

type Adventurer struct {
	ID       uuid.UUID
	World    uuid.UUID
	Hometown uuid.UUID
	Party    uuid.UUID
	Name     string
	Stock    struct {
		Name                string
		ChosenLevelBenefits []string
	}
	Raiment
	Condition
	Goals         AdventurerGoals
	Rewards       AdventurerRewards
	Relationships AdventurerRelationships
	Abilities     AdventurerAbilities
	Traits        map[string]*AdventurerTrait
	Skills        map[string]*AdventurerSkill
	Wises         map[string]*AdventurerWise
	Inventory     Inventory
}
