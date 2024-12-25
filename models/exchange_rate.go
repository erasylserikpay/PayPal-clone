package models

type ExchangeRate struct {
    BaseCurrency    string  `json:"base_currency"`
    TargetCurrency  string  `json:"target_currency"`
    Rate            float64 `json:"rate"`
}
