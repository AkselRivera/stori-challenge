package domain

// UserBalance represents a user balance
// @Description User balance details including balance, total debits and total credits
// @Model UserBalance
type UserBalance struct {
	Balance      float64 `json:"balance" example:"25.00"`
	TotalDebits  float64 `json:"total_debits" example:"10.00"`
	TotalCredits float64 `json:"total_credits" example:"15.00"`
}
