package controllers

import (
    "encoding/json"
    "net/http"
    "paypal-clone/services"
)

func RequestPasswordReset(w http.ResponseWriter, r *http.Request) {
    type Request struct {
        Email string `json:"email"`
    }

    var req Request
    json.NewDecoder(r.Body).Decode(&req)

    err := services.RequestPasswordReset(req.Email)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"message": "Password reset link sent"})
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
    type ResetRequest struct {
        Token    string `json:"token"`
        Password string `json:"password"`
    }

    var req ResetRequest
    json.NewDecoder(r.Body).Decode(&req)

    err := services.ResetPassword(req.Token, req.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"message": "Password reset successful"})
}
