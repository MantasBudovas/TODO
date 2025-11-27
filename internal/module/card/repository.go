package card

import (
	"gorm.io/gorm"

	"app/internal/entity"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(card *entity.Card) error {
	return r.db.Create(card).Error
}

func (r *Repository) GetAll() ([]entity.Card, error) {
	var cards []entity.Card
	result := r.db.Find(&cards)
	return cards, result.Error
}

func (r *Repository) GetByID(id int) (*entity.Card, error) {
	var card entity.Card
	result := r.db.First(&card, id)
	return &card, result.Error
}

func (r *Repository) Update(card *entity.Card) error {
	return r.db.Model(card).Updates(card).Error
}
