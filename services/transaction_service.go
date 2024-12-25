package services

import (
    "paypal-clone/models"
    "paypal-clone/repositories"
)

func CreateTransaction(transaction models.Transaction) {
    repositories.SaveTransaction(transaction)
}
