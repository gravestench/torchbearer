package account

import (
	"github.com/google/uuid"
)

type Account struct {
	ID          uuid.UUID
	Username    string
	Email       string
	Password    string
	Adventurers []uuid.UUID
	Activated   bool // need to do one-time-passcode recovery after sign up
}
