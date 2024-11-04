package ports

import "github.com/AkselRivera/stori-challenge/migration-service/pkg/domain"

type EmailService interface {
	Send(emailData domain.EmailData) error
}
