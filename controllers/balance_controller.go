package controllers

import (
    "encoding/json"
    "net/http"
    "paypal-clone/services"
    "paypal-clone/models"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
    userID := r.Header.Get("email")
    balance := services.GetBalance(userID)
    json.NewEncoder(w).Encode(models.Balance{UserID: userID, Amount: balance})
}
