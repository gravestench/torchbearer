package models

import (
	"github.com/google/uuid"
)

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
}
