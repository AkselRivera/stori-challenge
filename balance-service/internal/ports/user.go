package ports

import (
	"time"

	"github.com/AkselRivera/stori-challenge/balance-service/internal/domain"
)

type UserService interface {
	GetBalance(userID int, startDate time.Time, endDate time.Time) (domain.UserBalance, error)
}

type UserRepository interface {
	GetUserTransactions(userID int, startDate time.Time, endDate time.Time) ([]domain.Transaction, error)
}
