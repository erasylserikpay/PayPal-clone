package controllers

import (
    "encoding/json"
    "net/http"
    "paypal-clone/services"
)

func ConvertCurrency(w http.ResponseWriter, r *http.Request) {
    type ConversionRequest struct {
        Amount        float64 `json:"amount"`
        BaseCurrency  string  `json:"base_currency"`
        TargetCurrency string  `json:"target_currency"`
    }

    var req ConversionRequest
    json.NewDecoder(r.Body).Decode(&req)

    convertedAmount, err := services.ConvertCurrency(req.Amount, req.BaseCurrency, req.TargetCurrency)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]float64{"converted_amount": convertedAmount})
}
