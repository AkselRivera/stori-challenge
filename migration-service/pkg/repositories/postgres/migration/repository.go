package migration

import (
	"github.com/AkselRivera/stori-challenge/migration-service/pkg/ports"
	"gorm.io/gorm"
)

var _ ports.MigrationRepository = &Repository{}

type Repository struct {
	Client *gorm.DB
}
