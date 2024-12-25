package models

type Transaction struct {
    ID       string  `json:"id"`
    Amount   float64 `json:"amount"`
    UserID   string  `json:"user_id"`
    Currency string  `json:"currency"`
}
