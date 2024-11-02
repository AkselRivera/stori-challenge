package migration

import (
	"strconv"
	"time"

	"github.com/AkselRivera/stori-challenge/migration-service/pkg/domain"
)

func (s *Service) ValidateData(data [][]string) ([]*domain.Transaction, error) {

	if len(data) == 0 {
		return nil, domain.HandleError(domain.ErrorInvalidDataType, "no data provided")
	}

	var transactions []*domain.Transaction
	for _, row := range data {

		var transaction domain.Transaction

		userID, err := strconv.Atoi(row[1])

		if err != nil {
			return nil, domain.HandleError(domain.ErrorInvalidDataType, "user_id must be an integer")
		}

		transaction.UserID = userID

		amount, err := strconv.ParseFloat(row[2], 64)

		if err != nil {
			return nil, domain.HandleError(domain.ErrorInvalidDataType, "amount must be a decimal number")
		}

		transaction.Amount = amount

		datetime, err := time.Parse(time.RFC3339, row[3])

		if err != nil {
			return nil, domain.HandleError(domain.ErrorInvalidDataType, "datetime must be a RFC3339 date")
		}

		transaction.DateTime = datetime

		transactions = append(transactions, &transaction)
	}

	return transactions, nil
}
