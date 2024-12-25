package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
    "log"
)

var DB *gorm.DB

func Connect() {
    var err error
    // Замените параметры подключения на ваши
    DB, err = gorm.Open("postgres", "host=localhost user=postgres dbname=paypalclone sslmode=disable password=Erasyl2007erasyl")
    if err != nil {
        log.Fatal(err)
    }
}
