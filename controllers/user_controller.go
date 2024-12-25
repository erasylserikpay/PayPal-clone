package controllers

import (
    "encoding/json"
    "net/http"
    "paypal-clone/models"
    "paypal-clone/services"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)
    services.CreateUser(user)
    json.NewEncoder(w).Encode(user)
}
