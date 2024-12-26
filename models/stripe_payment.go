package models

type StripePayment struct {
    Amount   int64  `json:"amount"`
    Currency string `json:"currency"`
    Source   string `json:"source"`
}
