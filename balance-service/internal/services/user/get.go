package user

import (
	"math"
	"time"

	"github.com/AkselRivera/stori-challenge/balance-service/internal/domain"
)

func (s *Service) GetBalance(userID int, startDate time.Time, endDate time.Time) (domain.UserBalance, error) {

	transactions, err := s.Repo.GetUserTransactions(userID, startDate, endDate)

	if err != nil {
		return domain.UserBalance{}, domain.HandleError(err, "")
	}

	var balance domain.UserBalance

	for _, transaction := range transactions {
		if transaction.Amount > 0 {
			balance.TotalCredits += transaction.Amount
		} else {
			balance.TotalDebits += -transaction.Amount
		}
	}

	balance.TotalCredits = math.Floor(balance.TotalCredits*100) / 100
	balance.TotalDebits = math.Floor(balance.TotalDebits*100) / 100
	balance.Balance = math.Floor((balance.TotalCredits+balance.TotalDebits)*100) / 100

	return balance, nil
}
