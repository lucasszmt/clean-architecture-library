package migrations

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	Id   uint `gorm:"primaryKey"`
	ISBN string
	//Title       string
	Author      string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
