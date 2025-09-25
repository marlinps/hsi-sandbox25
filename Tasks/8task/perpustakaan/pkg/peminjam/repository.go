package peminjam

import (
	"perpustakaan_app/pkg/entities"

	"gorm.io/gorm"
)

type PeminjamRepository interface {
	CreatePeminjam(peminjam *entities.Peminjam) (*entities.Peminjam, error)
	ReadPeminjam() ([]entities.Peminjam, error)
	UpdatePeminjam(peminjam *entities.Peminjam) (*entities.Peminjam, error)
	DeletePeminjam(ID uint) error
}

type peminjamRepository struct {
	db *gorm.DB
}

func NewPeminjamRepository(db *gorm.DB) PeminjamRepository {
	return &peminjamRepository{
		db: db,
	}
}

func (r *peminjamRepository) CreatePeminjam(peminjam *entities.Peminjam) (*entities.Peminjam, error) {
	if err := r.db.Create(peminjam).Error; err != nil {
		return nil, err
	}
	return peminjam, nil
}

func (r *peminjamRepository) ReadPeminjam() ([]entities.Peminjam, error) {
	var peminjams []entities.Peminjam
	if err := r.db.Model(&entities.Peminjam{}).Preload("BukuDiPinjam").Find(&peminjams).Error; err != nil {
		return nil, err
	}
	return peminjams, nil
}

func (r *peminjamRepository) UpdatePeminjam(peminjam *entities.Peminjam) (*entities.Peminjam, error) {
	if err := r.db.Save(peminjam).Error; err != nil {
		return nil, err
	}
	return peminjam, nil
}

func (r *peminjamRepository) DeletePeminjam(ID uint) error {
	if err := r.db.Delete(&entities.Peminjam{}, ID).Error; err != nil {
		return err
	}
	return nil
}
