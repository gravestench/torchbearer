package account

import (
	"encoding/json"

	"github.com/google/uuid"

	"torchbearer/pkg/services/config"
)

func (s *Service) ConfigFileName() string {
	return "accounts.json"
}

func (s *Service) LoadConfig(config *config.Config) {
	if s.accounts == nil {
		s.accounts = make(map[uuid.UUID]Account)
	}

	for _, key := range config.GroupKeys() {
		data, err := config.Group(key).Marshal()
		if err != nil {
			s.logger.Error("marshalling account data", "error", err)
			continue
		}

		var a Account

		err = json.Unmarshal(data, &a)
		if err != nil {
			s.logger.Error("unmarshalling account data", "error", err)
			continue
		}

		s.accounts[a.ID] = a
	}
}
