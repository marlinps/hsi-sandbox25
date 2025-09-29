package buku

import (
	"errors"

	"perpustakaan_app/pkg/entities"
)

type BukuService interface {
	InsertBuku(buku *entities.Buku) (*entities.Buku, error)
	FetchBuku() ([]entities.Buku, error)
	UpdateBuku(buku *entities.Buku) (*entities.Buku, error)
	RemoveBuku(id uint) error
}

type bukuService struct {
	bukuRepository BukuRepository
}

func NewBukuService(b BukuRepository) BukuService {
	return &bukuService{
		bukuRepository: b,
	}
}

func (s *bukuService) InsertBuku(buku *entities.Buku) (*entities.Buku, error) {
	if buku.Judul == "" {
		return nil, errors.New("judul buku tidak boleh kosong")
	}
	return s.bukuRepository.CreateBuku(buku)
}

func (s *bukuService) FetchBuku() ([]entities.Buku, error) {
	return s.bukuRepository.ReadBuku()
}

func (s *bukuService) UpdateBuku(buku *entities.Buku) (*entities.Buku, error) {
	if buku.ID == 0 {
		return nil, errors.New("ID buku tidak boleh kosong")
	}
	return s.bukuRepository.UpdateBuku(buku)
}

func (s *bukuService) RemoveBuku(id uint) error {
	if id == 0 {
		return errors.New("ID buku tidak boleh kosong")
	}
	return s.bukuRepository.DeleteBuku(id)
}
