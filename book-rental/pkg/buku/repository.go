package buku

import (
	"book_rental_app/pkg/entities"

	"gorm.io/gorm"
)

type BukuRepository struct {
	db *gorm.DB
}

func NewBukuRepository(db *gorm.DB) *BukuRepository {
	return &BukuRepository{
		db: db,
	}
}

func (r *BukuRepository) Create(buku entities.Buku) error {
	return r.db.Create(&buku).Error
}

func (r *BukuRepository) GetAll() ([]entities.Buku, error) {
	var buku []entities.Buku
	err := r.db.Find(&buku).Error
	return buku, err
}
