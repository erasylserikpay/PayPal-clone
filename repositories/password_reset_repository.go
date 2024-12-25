package repositories

import (
    "paypal-clone/models"
    "paypal-clone/db"
)

func SavePasswordResetRequest(resetRequest models.PasswordReset) {
    db.DB.Create(&resetRequest)
}

func FindPasswordResetRequestByToken(token string) (models.PasswordReset, error) {
    var resetRequest models.PasswordReset
    if err := db.DB.Where("token = ?", token).First(&resetRequest).Error; err != nil {
        return resetRequest, err
    }
    return resetRequest, nil
}

func DeletePasswordResetRequest(token string) {
    db.DB.Where("token = ?", token).Delete(&models.PasswordReset{})
}
