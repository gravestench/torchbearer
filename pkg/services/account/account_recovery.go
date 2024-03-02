package account

import (
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) SendOneTimePassCode(email string) error {
	account, err := s.GetAccountByEmail(email)
	if err != nil {
		return fmt.Errorf("account with email %q not found", email)
	}

	passcode := uuid.New().String()

	if s.accountRecoveryPasscodes == nil {
		s.accountRecoveryPasscodes = make(map[uuid.UUID]string)
	}

	s.accountRecoveryPasscodes[account.ID] = passcode

	if err = s.sendOneTimePassCodeEmail(email, passcode); err != nil {
		return err
	}

	return nil
}

func (s *Service) sendOneTimePassCodeEmail(email, code string) error {
	if s.email == nil {
		return fmt.Errorf("cannot send email: no email service found")
	}

	const (
		sender                    = "torchbearer@gmail.com"
		subject                   = "Account Recovery"
		fmtAccountRecoveryMessage = `Your account recovery passcode is: %s`
	)

	return s.email.SendEmail(sender, email, subject, fmt.Sprintf(fmtAccountRecoveryMessage, code))
}
