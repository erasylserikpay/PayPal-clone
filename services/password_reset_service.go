package services

import (
    "crypto/rand"
    "encoding/hex"
    "errors"
    "log"
    "paypal-clone/models"
    "paypal-clone/repositories"
    "paypal-clone/utils"
    "time"
)

func generateResetToken() string {
    b := make([]byte, 16)
    rand.Read(b)
    return hex.EncodeToString(b)
}

func RequestPasswordReset(email string) error {
    _, err := repositories.FindUserByEmail(email)
    if err != nil {
        log.Println("User not found:", err)
        return errors.New("user not found")
    }

    token := generateResetToken()
    expiry := time.Now().Add(1 * time.Hour)

    resetRequest := models.PasswordReset{
        Email:  email,
        Token:  token,
        Expiry: expiry,
    }

    repositories.SavePasswordResetRequest(resetRequest)

    resetLink := "http://localhost:8080/reset-password?token=" + token
    emailBody := "Для сброса пароля перейдите по следующей ссылке: " + resetLink
    return utils.SendEmail(email, "Запрос на сброс пароля", emailBody)
}

func ResetPassword(token string, newPassword string) error {
    resetRequest, err := repositories.FindPasswordResetRequestByToken(token)
    if err != nil {
        log.Println("Token not found or expired:", err)
        return errors.New("invalid or expired token")
    }
    if time.Now().After(resetRequest.Expiry) {
        log.Println("Token expired")
        return errors.New("invalid or expired token")
    }

    user, err := repositories.FindUserByEmail(resetRequest.Email)
    if err != nil {
        log.Println("User not found:", err)
        return errors.New("user not found")
    }

    user.Password = newPassword
    repositories.UpdateUser(user)
    repositories.DeletePasswordResetRequest(token)

    return nil
}
