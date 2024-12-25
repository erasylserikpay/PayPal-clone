package models

import "time"

type PasswordReset struct {
    Email       string    `json:"email"`
    Token       string    `json:"token"`
    Expiry      time.Time `json:"expiry"`
}
