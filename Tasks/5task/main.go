package main

import (
	"5task/pegawai"
	"fmt"
)

func main() {
	p := pegawai.Pegawai{Nama: "John Doe",
		Posisi:      "Manager",
		GajiBulanan: 100000000.0,
	}

	fmt.Printf("Gaji Tahunan: %.2f\n", p.HitungGajiTahunan())

	p.TampilkanInformasi()

}
