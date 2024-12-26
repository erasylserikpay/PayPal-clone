package controllers

import (
    "encoding/json"
    "net/http"
    "paypal-clone/models"
    "paypal-clone/services"
)

func CreatePayment(w http.ResponseWriter, r *http.Request) {
    var payment models.StripePayment
    json.NewDecoder(r.Body).Decode(&payment)

    charge, err := services.ProcessPayment(payment)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(charge)
}
