package main

import "fmt"

type Pegawai struct {
	Nama        string
	Posisi      string
	GajiBulanan float64
}

type InformasiPegawai interface {
	TampilkanInformasi()
}

func (p Pegawai) HitungGajiTahunan() float64 {
	return p.GajiBulanan * 12
}

func (p Pegawai) TampilkanInformasi() {
	fmt.Printf("Nama: %s\n", p.Nama)
	fmt.Printf("Posisi: %s\n", p.Posisi)
	fmt.Printf("Gaji Bulanan: %.2f\n", p.GajiBulanan)
}
