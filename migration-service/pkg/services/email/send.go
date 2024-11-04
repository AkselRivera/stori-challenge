package email

import (
	"github.com/AkselRivera/stori-challenge/migration-service/pkg/domain"
)

func (s *Service) Send(email domain.EmailData) error {
	return s.Sender.Send(email)
}
