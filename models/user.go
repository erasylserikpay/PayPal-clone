package models

import "time"

type User struct {
    ID                string    `json:"id"`
    Name              string    `json:"name"`
    Email             string    `json:"email"`
    Password          string    `json:"password"`
    Verified          bool      `json:"verified"`
    VerificationCode  string    `json:"-"`
    VerificationExpiry time.Time `json:"-"`
}
