package user

import (
	"time"

	"github.com/AkselRivera/stori-challenge/balance-service/internal/domain"
)

func (r *Repository) GetUserTransactions(userID int, startDate time.Time, endDate time.Time) ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	var count int64

	db := r.Client.Model(&domain.Transaction{}).Where("user_id = ?", userID).Count(&count)

	if db.Error != nil {
		return nil, db.Error
	}

	if count == 0 {
		return nil, domain.ErrorUserNotFound
	}

	query := "user_id = ? "

	if !startDate.IsZero() && !endDate.IsZero() {
		query += "AND date_time BETWEEN ? AND ? "
		if db = r.Client.Where(query, userID, startDate, endDate).Find(&transactions); db.Error != nil {
			return nil, db.Error
		}

	} else {
		if db = r.Client.Where(query, userID).Find(&transactions); db.Error != nil {
			return nil, db.Error
		}
	}

	return transactions, nil
}
