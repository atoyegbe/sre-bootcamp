package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "github.com/atoyegbe/sre-bootcamp/models"
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database")
    }

    // Migrate the schema
    DB.AutoMigrate(&models.Student{})
}
