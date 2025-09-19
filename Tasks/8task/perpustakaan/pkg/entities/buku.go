package entities

import "time"

type Buku struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Judul     string    `json:"judul"`
	Tahun     int       `json:"tahun"`
	Penulis   string    `json:"penulis"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
