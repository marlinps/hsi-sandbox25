package main

import (
	"5task/pegawai"
)

func main() {
	p := pegawai.Pegawai{Nama: "John Doe",
		Posisi:      "Manager",
		GajiBulanan: 30000000.00,
	}

	p.TampilkanInformasi()
}
