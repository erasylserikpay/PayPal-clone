package repositories

import (
    "paypal-clone/models"
    "paypal-clone/db"
)

func SaveUser(user models.User) {
    db.DB.Create(&user)
}

func UpdateUser(user models.User) {
    db.DB.Save(&user)
}

func FindUserByEmail(email string) (models.User, error) {
    var user models.User
    if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
        return user, err
    }
    return user, nil
}
