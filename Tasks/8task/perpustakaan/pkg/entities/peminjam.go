package entities

import (
	"time"

	"gorm.io/gorm"
)

type Peminjam struct {
	gorm.Model
	TanggalPeminjaman   time.Time
	TanggalPengembalian time.Time // dihitung berdasarkan tanggal peminjaman + 7 hari
	NamaPeminjam        string
	BukuDiPinjam        []Buku
}

func (p *Peminjam) BeforeCreate(tx *gorm.DB) (err error) {
	if p.TanggalPeminjaman.IsZero() {
		p.TanggalPeminjaman = time.Now()
	}

	p.TanggalPengembalian = p.TanggalPeminjaman.AddDate(0, 0, 7)
	return
}
