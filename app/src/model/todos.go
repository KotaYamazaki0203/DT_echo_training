package model

import (
	"gorm.io/gorm"
	"time"
)

type Todos struct {
	ID        uint
	TITLE     string    `gorm:"column:title;type:varchar(255);not null"`
	CONTENT   string    `gorm:"column:content;not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt gorm.DeletedAt
}
