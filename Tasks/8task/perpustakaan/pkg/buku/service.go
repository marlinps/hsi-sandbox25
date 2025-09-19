package buku

import (
	"errors"
	"perpustakaan_app/pkg/entities"
)

// Kontrak Service
type BukuService interface {
	CreateBuku(judul string, penulis string, tahun int) (*entities.Buku, error)
	GetAllBuku() ([]entities.Buku, error)
	GetBukuByID(ID uint) (*entities.Buku, error)
	UpdateBuku(ID uint, judul string, penulis string, tahun int) (*entities.Buku, error)
	DeleteBuku(ID uint) error
}

// Implementasi Service
type bukuService struct {
	repo BukuRepository
}

func NewBukuService(repo BukuRepository) BukuService {
	return &bukuService{repo: repo}
}

func (s *bukuService) CreateBuku(judul string, penulis string, tahun int) (*entities.Buku, error) {
	if judul == "" {
		return nil, errors.New("judul tidak boleh kosong")
	}
	if penulis == "" {
		return nil, errors.New("penulis tidak boleh kosong")
	}

	buku := &entities.Buku{
		Judul:   judul,
		Penulis: penulis,
		Tahun:   tahun,
	}
	return s.repo.CreateBuku(buku)
}

func (s *bukuService) GetAllBuku() ([]entities.Buku, error) {
	return s.repo.ReadBuku()
}

func (s *bukuService) GetBukuByID(ID uint) (*entities.Buku, error) {
	return s.repo.ReadBukuByID(ID)
}

func (s *bukuService) UpdateBuku(ID uint, judul string, penulis string, tahun int) (*entities.Buku, error) {
	buku, err := s.repo.ReadBukuByID(ID)
	if err != nil {
		return nil, err
	}

	// Update field jika ada perubahan
	if judul != "" {
		buku.Judul = judul
	}
	if penulis != "" {
		buku.Penulis = penulis
	}
	if tahun != 0 {
		buku.Tahun = tahun
	}

	return s.repo.UpdateBuku(buku)
}

func (s *bukuService) DeleteBuku(ID uint) error {
	return s.repo.DeleteBuku(ID)
}
