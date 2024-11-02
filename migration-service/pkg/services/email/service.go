package email

import "github.com/AkselRivera/stori-challenge/migration-service/pkg/ports"

var _ ports.EmailService = &Service{}

type Service struct {
}
