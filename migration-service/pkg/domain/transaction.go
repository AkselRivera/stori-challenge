package domain

import "time"

type Transaction struct {
	ID       int
	UserID   int
	Amount   float64
	DateTime time.Time
}

var ValidColumns = []string{"id", "user_id", "amount", "datetime"}
