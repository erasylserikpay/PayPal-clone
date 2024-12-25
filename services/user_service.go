package services

import (
    "paypal-clone/models"
    "paypal-clone/repositories"
)

func CreateUser(user models.User) {
    repositories.SaveUser(user)
}
