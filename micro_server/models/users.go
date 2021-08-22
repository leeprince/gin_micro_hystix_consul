package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Id         int64 `gorm:"primaryKey"`
	Username   string
	Password   string
	Age        int
	CreatedAt  int64
	UpdatedAt  int64
	Deleted_at int64
}

func (m Users) TableName() string {
	return "users"
}
