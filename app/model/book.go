package model

import (
	"com.github.gin-common/common/models"
)

type Book struct {
	models.BaseModel
	models.SoftDeleteModel
	BookName string  `gorm:"not null;size:256" json:"book_name"`
	Price    float64 `gorm:"not null;precision:6;scale:2" json:"price"`
	UserID   uint    `gorm:"not null" json:"user_id"`
}
