package migration

import (
	"fmt"

	"github.com/AkselRivera/stori-challenge/migration-service/pkg/domain"
)

func (s *Service) ValidateFileHeaders(data []string) error {

	if len(data) < len(domain.ValidColumns) {
		return domain.HandleError(domain.ErrorInvalidCsvColumns, fmt.Sprintf("invalid number of headers got: '%d' expected: '%d'", len(data), len(domain.ValidColumns)))
	}

	for index, header := range data {
		if header != domain.ValidColumns[index] {
			errDetails := fmt.Sprintf("invalid header got: '%s' expected: '%s'", header, domain.ValidColumns[index])
			return domain.HandleError(domain.ErrorInvalidCsvColumns, errDetails)
		}

		if index+1 == len(domain.ValidColumns) {
			break
		}
	}

	return nil
}
