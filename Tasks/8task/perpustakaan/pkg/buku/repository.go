package buku

import (
	"perpustakaan_app/pkg/entities"

	"gorm.io/gorm"
)

type BukuRepository interface {
	CreateBuku(buku *entities.Buku) (*entities.Buku, error)
	GetAllBuku() ([]entities.Buku, error)
	UpdateBuku(buku *entities.Buku) (*entities.Buku, error)
	DeleteBuku(ID uint) error
}

type bukuRepository struct {
	db *gorm.DB
}

func NewBukuRepository(db *gorm.DB) BukuRepository {
	return &bukuRepository{
		db: db,
	}
}

func (r *bukuRepository) CreateBuku(buku *entities.Buku) (*entities.Buku, error) {
	if err := r.db.Create(buku).Error; err != nil {
		return nil, err
	}
	return buku, nil
}

func (r *bukuRepository) GetAllBuku() ([]entities.Buku, error) {
	var bukus []entities.Buku
	if err := r.db.Find(&bukus).Error; err != nil {
		return nil, err
	}
	return bukus, nil
}

func (r *bukuRepository) UpdateBuku(buku *entities.Buku) (*entities.Buku, error) {
	if err := r.db.Save(buku).Error; err != nil {
		return nil, err
	}
	return buku, nil
}

func (r *bukuRepository) DeleteBuku(ID uint) error {
	return r.db.Delete(&entities.Buku{}, ID).Error
}
