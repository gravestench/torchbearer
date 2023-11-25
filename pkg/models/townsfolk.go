package models

import (
	"github.com/google/uuid"
)

type Townsfolk struct {
	ID            uuid.UUID
	World         uuid.UUID
	Hometown      uuid.UUID
	Residence     uuid.UUID
	Relationships map[uuid.UUID]string
	Skills        []Record
	Name          string
	Notes         []string
}
