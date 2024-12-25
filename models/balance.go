package models

type Balance struct {
    UserID string  `json:"user_id"`
    Amount float64 `json:"amount"`
}
