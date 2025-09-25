package peminjam

import (
	"errors"
	"perpustakaan_app/pkg/entities"
)

type PeminjamService interface {
	InsertPeminjam(peminjam *entities.Peminjam) (*entities.Peminjam, error)
	FetchPeminjam() ([]entities.Peminjam, error)
	UpdatePeminjam(peminjam *entities.Peminjam) (*entities.Peminjam, error)
	RemovePeminjam(ID uint) error
}

type peminjamService struct {
	peminjamRepository PeminjamRepository
}

func NewPeminjamService(r PeminjamRepository) PeminjamService {
	return &peminjamService{
		peminjamRepository: r,
	}
}

func (s *peminjamService) InsertPeminjam(peminjam *entities.Peminjam) (*entities.Peminjam, error) {
	return s.peminjamRepository.CreatePeminjam(peminjam)
}

func (s *peminjamService) FetchPeminjam() ([]entities.Peminjam, error) {
	return s.peminjamRepository.ReadPeminjam()
}

func (s *peminjamService) UpdatePeminjam(peminjam *entities.Peminjam) (*entities.Peminjam, error) {
	if peminjam.ID == 0 {
		return nil, errors.New("ID peminjam tidak boleh kosong")
	}
	return s.peminjamRepository.UpdatePeminjam(peminjam)
}

func (s *peminjamService) RemovePeminjam(ID uint) error {
	if ID == 0 {
		return errors.New("ID peminjam tidak boleh kosong")
	}
	return s.peminjamRepository.DeletePeminjam(ID)
}
