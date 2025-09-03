package pegawai

import "fmt"

type Pegawai struct {
	ID          uint   `gorm:"primaryKey"`
	Nama        string `gorm:"size:255"`
	Posisi      string `gorm:"size:100"`
	GajiBulanan float64
}

type InformasiPegawai interface {
	TampilkanInformasi()
}

func (p Pegawai) HitungGajiTahunan() float64 {
	return 12 * p.GajiBulanan
}

func (p Pegawai) TampilkanInformasi() {
	fmt.Printf("Nama: %s\n", p.Nama)
	fmt.Printf("Posisi: %s\n", p.Posisi)
	fmt.Printf("Gaji Tahunan: %.2f\n", p.HitungGajiTahunan())
}
