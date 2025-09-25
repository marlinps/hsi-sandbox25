package entities

import "gorm.io/gorm"

type Buku struct {
	gorm.Model `json:"id"`
	Judul      string `json:"judul"`
	Tahun      int    `json:"tahun"`
	Penulis    string `json:"penulis"`
	PeminjamID uint   `json:"peminjam_id"`
}
