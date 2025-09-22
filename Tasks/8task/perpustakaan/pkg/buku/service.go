package buku

import (
	"errors"
	"perpustakaan_app/pkg/entities"
)

type BukuRepo interface {
	InsertBuku(buku *entities.Buku) (*entities.Buku, error)
	GetAllBuku() ([]entities.Buku, error)
	UpdateBuku(buku *entities.Buku) (*entities.Buku, error)
	DeleteBuku(id uint) error
}

type BukuService struct {
	repo BukuRepo
}

func NewBukuService(repo BukuRepo) *BukuService {
	return &BukuService{
		repo: repo,
	}
}

func (s *BukuService) InsertBuku(buku *entities.Buku) (*entities.Buku, error) {
	// Validasi sederhana
	if buku.Judul == "" {
		return nil, errors.New("judul buku tidak boleh kosong")
	}
	return s.repo.InsertBuku(buku)
}

func (s *BukuService) GetAllBuku() ([]entities.Buku, error) {
	return s.repo.GetAllBuku()
}

func (s *BukuService) UpdateBuku(buku *entities.Buku) (*entities.Buku, error) {
	if buku.ID == 0 {
		return nil, errors.New("ID buku tidak boleh kosong")
	}
	return s.repo.UpdateBuku(buku)
}

func (s *BukuService) DeleteBuku(id uint) error {
	if id == 0 {
		return errors.New("ID buku tidak boleh kosong")
	}
	return s.repo.DeleteBuku(id)
}
