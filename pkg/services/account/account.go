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
}
