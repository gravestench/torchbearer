package models

import (
	"github.com/google/uuid"
)

type AdventurerRelationships struct {
	Parents   []uuid.UUID
	Mentor    uuid.UUID
	Friend    uuid.UUID
	Enemies   []uuid.UUID
	Allies    []uuid.UUID
	Anecdotes []string
}
