package database

import (
	`time`

	`gorm.io/gorm`
)

type BaseModel struct {
	ID        string         `json:"id" gorm:"primarykey;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
