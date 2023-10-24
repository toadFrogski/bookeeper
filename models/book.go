package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `gorm:"->:false;column:created_at" json:"-"`
	UpdatedAt time.Time      `gorm:"->:false;column:updated_at" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"->:false;column:deleted_at" json:"-"`
}

type Book struct {
	ID          int    `gorm:"column:id; primary_key; not null; autoIncrement" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Author      string `gorm:"column:author" json:"author"`
	Description string `gorm:"colunm:description; text" json:"description"`
	BaseModel
}
