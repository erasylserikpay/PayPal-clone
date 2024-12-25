package services

import (
    "crypto/rand"
    "encoding/hex" 
    "errors"
    "paypal-clone/models"
    "github.com/dgrijalva/jwt-go"
    "paypal-clone/repositories" 
    "time"
)


func generateVerificationCode() string {
    b := make([]byte, 6)
    rand.Read(b)
    return hex.EncodeToString(b)
}
var jwtKey = []byte("your_secret_key")

func RegisterUser(user models.User) string { 
    verificationCode := generateVerificationCode() 
    user.VerificationCode = verificationCode 
    user.VerificationExpiry = time.Now().Add(24 * time.Hour) 
    repositories.SaveUser(user) 
    return verificationCode
}

func LoginUser(credentials models.Credentials) (string, error) {
    user, err := repositories.FindUserByEmail(credentials.Email)
    if err != nil || user.Password != credentials.Password {
        return "", errors.New("invalid credentials")
    }
    
    // Create JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": user.Email,
        "exp":   time.Now().Add(time.Hour * 72).Unix(),
    })
    
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }
    
    return tokenString, nil
}
func VerifyUser(email string, code string) error { 
    user, err := repositories.FindUserByEmail(email) 
        if err != nil || user.VerificationCode != code || time.Now().After(user.VerificationExpiry) { 
            return errors.New("invalid or expired verification code") } 
            user.Verified = true 
            user.VerificationCode = "" 
            repositories.UpdateUser(user) 
            return nil
 }