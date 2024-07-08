package models

type Student struct {
    ID    uint   `json:"id" gorm:"primaryKey"`
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}