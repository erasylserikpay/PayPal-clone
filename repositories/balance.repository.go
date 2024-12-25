package repositories

import (
    "paypal-clone/models"
    "paypal-clone/db"
)

func GetUserBalance(userID string) float64 {
    var balance models.Balance
    db.DB.Where("user_id = ?", userID).First(&balance)
    return balance.Amount
}
