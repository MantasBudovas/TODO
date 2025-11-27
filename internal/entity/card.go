package entity

import (
	"app/pkg/entity"
	"app/pkg/types"
)

type Card struct {
	entity.Entity
	Name      string      `json:"name"`
	Completed bool        `json:"completed"`
	DueDate   *types.Date `json:"due_date" gorm:"type:datetime"`
	Priority  string      `json:"priority"`
}
