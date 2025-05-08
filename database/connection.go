package database

import (
    "log"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "github.com/K0725/task-tracker/models"
)

var DB *gorm.DB

func ConnectDB() {
    db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    db.AutoMigrate(&models.Task{})
    DB = db
}
