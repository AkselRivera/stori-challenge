package migration

import "github.com/AkselRivera/stori-challenge/migration-service/pkg/ports"

var _ ports.MigrationService = &Service{}

type Service struct {
	Repo   ports.MigrationRepository
	Sender ports.EmailService
}
