package services

import (
    "paypal-clone/repositories"
)

func GetBalance(userID string) float64 {
    return repositories.GetUserBalance(userID)
}
