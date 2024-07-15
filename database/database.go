package database

import (
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "github.com/atoyegbe/sre-bootcamp/models"
)

var DB *gorm.DB

func Connect() error {
    var err error
    DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        return err
    }

    if err := DB.AutoMigrate(&models.Student{}); err != nil {
        return err
    }

    log.Println("Connected to database and migrated schema")
    return nil
}

func Close() {
    if DB != nil {
        sqlDB, err := DB.DB()
        if err != nil {
            log.Printf("Error getting underlying SQL DB: %v", err)
            return
        }
        if err := sqlDB.Close(); err != nil {
            log.Printf("Error closing database connection: %v", err)
        }
    }
}