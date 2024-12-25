package repositories

import (
    "paypal-clone/models"
    "paypal-clone/db"
)

func SaveTransaction(transaction models.Transaction) {
    db.DB.Create(&transaction)
}
