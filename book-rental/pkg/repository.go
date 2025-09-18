package buku

import (
	"book_rental_app/pkg/entities"
)

type Repository interface {
	GetAll() ([]entities.Buku, error)
	GetByID(id int) (entities.Buku, error)
	Create(buku entities.Buku) (entities.Buku, error)
}
