package email

import (
	"crypto/tls"
	"fmt"
	"log/slog"

	"github.com/gravestench/servicemesh"
	"gopkg.in/gomail.v2"

	"torchbearer/pkg/services/config"
)

type Service struct {
	logger     *slog.Logger
	cfgManager config.Dependency
}

func (s *Service) SendEmail(from, to, subject, body string) error {
	// Set up the email sender using SMTP
	cfg, err := s.cfgManager.GetConfigByFileName(s.ConfigFileName())
	if err != nil {
		return fmt.Errorf("loading config: %v", err)
	}

	authUser := cfg.Group("auth").GetString("user")
	authPass := cfg.Group("auth").GetString("password")

	// Create a new message
	message := gomail.NewMessage()
	message.SetHeader("From", from) // Replace with your email address
	message.SetHeader("To", to)     // Replace with the recipient's email address
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, authUser, authPass) // Replace with your credentials
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email
	if err = d.DialAndSend(message); err != nil {
		return fmt.Errorf("sending email: %v", err)
	}

	return nil
}

func (s *Service) Init(mesh servicemesh.Mesh) {
}

func (s *Service) Name() string {
	return "Email"
}

func (s *Service) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *slog.Logger {
	return s.logger
}
