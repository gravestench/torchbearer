package session

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID                uuid.UUID
	PreviousSessionID uuid.UUID
	Log               map[time.Time]string
	CurrentPhase      string
}
