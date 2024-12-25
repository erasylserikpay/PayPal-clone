package controllers

import (
    "encoding/json"
    "net/http"
    "paypal-clone/models"
    "paypal-clone/services"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
    var transaction models.Transaction
    json.NewDecoder(r.Body).Decode(&transaction)
    services.CreateTransaction(transaction)
    json.NewEncoder(w).Encode(transaction)
}
