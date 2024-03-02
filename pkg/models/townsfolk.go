package models

import (
	"github.com/google/uuid"
)

func NewTownsfolk() *Townsfolk {
	return &Townsfolk{
		TownsfolkID:   uuid.UUID{},
		Relationships: make(map[uuid.UUID]string),
		Skills:        make([]SkillRecord, 0),
		Notes:         make([]string, 0),
	}
}

type Townsfolk struct {
	TownsfolkID   uuid.UUID
	WorldID       uuid.UUID
	HometownID    uuid.UUID
	ResidenceID   uuid.UUID
	Relationships map[uuid.UUID]string
	Skills        []SkillRecord
	Name          string
	Notes         []string
}
