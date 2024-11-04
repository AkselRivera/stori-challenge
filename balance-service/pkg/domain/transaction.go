package domain

import "time"

// Transaction represents an user transaction
// Description Transaction details including id, user_id, amount and datetime
// Model Transaction
type Transaction struct {
	ID       int
	UserID   int
	Amount   float64
	DateTime time.Time
}
