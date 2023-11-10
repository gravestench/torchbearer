package models

import (
	"github.com/google/uuid"
)

type Adventurer struct {
	ID   uuid.UUID
	Name string
	Age  int
	Stock
	Traits []AdventurerTrait
	Wises
	Raiment
	Condition
	Goals         AdventurerGoals
	Rewards       AdventurerRewards
	Relationships AdventurerRelationships
	Abilities     AdventurerAbilities
	Skills        []AdventurerSkill
}
