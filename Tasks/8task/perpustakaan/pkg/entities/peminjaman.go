package entities

import "time"

type Peminjaman struct {
	ID                  uint `gorm:"primaryKey"`
	TanggalPeminjaman   time.Time
	TanggalPengembalian time.Time // dihitung berdasarkan tanggal peminjaman + 7 hari
	NamaPeminjam        string
	NamaBuku            string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
