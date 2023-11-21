package models

import (
	"github.com/google/uuid"
)

type Actor struct {
	ID       uuid.UUID
	World    uuid.UUID
	Hometown uuid.UUID
	Name     string
}
