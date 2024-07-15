package models

import "time"

type Student struct {
    ID    uint   `json:"id" gorm:"primaryKey"`
    Name  string `json:"name"`
    Age   uint8    `json:"age"`
    Email string `json:"email"`
    CreatedAt time.Time `json:"created_at"` // Automatically managed by GORM for creation time
    UpdatedAt time.Time `json:"updated_at"`
}