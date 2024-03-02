package email

import (
	"torchbearer/pkg/services/config"
)

type Dependency = EmailSender

type EmailSender interface {
	config.HasDefaultConfig
	SendEmail(from, to, subject, body string) error
}
