package migration

import (
	"github.com/AkselRivera/stori-challenge/migration-service/pkg/ports"
)

type Handler struct {
	MigrationService ports.MigrationService
}
