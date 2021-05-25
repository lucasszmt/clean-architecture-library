package migrations

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	Phone     string
	Cpf       string
	Address   string
	Teste     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
