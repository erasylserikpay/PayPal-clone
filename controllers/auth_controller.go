package controllers

import (
    "encoding/json"
    "net/http"
    "paypal-clone/models"
    "paypal-clone/services"
    "paypal-clone/utils"
    //"time"
)

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)
    verificationCode := services.RegisterUser(user)

    // Send verification email
    emailBody := "Ваш код подтверждения: " + verificationCode
    utils.SendEmail(user.Email, "Подтверждение регистрации", emailBody)

    json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
    var credentials models.Credentials
    json.NewDecoder(r.Body).Decode(&credentials)
    token, err := services.LoginUser(credentials)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func VerifyEmail(w http.ResponseWriter, r *http.Request) {
    type VerificationRequest struct {
        Email string `json:"email"`
        Code  string `json:"code"`
    }

    var req VerificationRequest
    json.NewDecoder(r.Body).Decode(&req)

    err := services.VerifyUser(req.Email, req.Code)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"message": "Email verified successfully"})
}
