package ports

import "github.com/AkselRivera/stori-challenge/migration-service/pkg/domain"

type MigrationService interface {
	ValidateData(data [][]string) ([]*domain.Transaction, error)
	ValidateFileHeaders(data []string) error
	Migrate(transactions []*domain.Transaction, fileName string) (*domain.EmailData, error)
}

type MigrationRepository interface {
	InsertMany(t []*domain.Transaction) error
}
