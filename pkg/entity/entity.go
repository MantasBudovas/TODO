package entity

import (
	"time"
)

type Entity struct {
	ID        int       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `gorm:"index"       json:"created_at"`
	UpdatedAt time.Time `gorm:"index"       json:"updated_at"`
}
