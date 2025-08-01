package main

import "fmt"

func hitungLuasLingkaran(jariJari float32) float32 { // TODO: tipe data pengembalian ada dibelakang
	const phi float32 = 3.14
	var luas float32
	luas = phi * jariJari * jariJari
	return luas
}

func cetakHasilPerhitungan(hasilHitung float32) {
	fmt.Println("Luas Lingkaran: ", hasilHitung)
}

func main() {
	var usiaSaya int = 40
	// var phi float32 = 3.14

	// var test = "abc"
	// var apakahSayaBenar bool = true

	if usiaSaya < 5 {
		fmt.Println("Ini balita")
	} else {
		fmt.Println("Ini bukan lagi balita")
	}

	//perulangan
	for i := 1; i < 10; i++ {
		fmt.Println("Perulangan ke-", i)
	}

	const phi float32 = 3.14
	fmt.Println("Phi :", phi)

	//function
	var jariJari float32 = 17
	luasL := hitungLuasLingkaran(jariJari)
	cetakHasilPerhitungan(luasL)
}

/* TODO: standart penulisan di Golang menggunakan camelCase
TODO: Reminder dalam memilih tipe data jangan asal-asal harus dipertimbangkan sesuai kebutuhan, alokasi memory;server dsb
TODO: Flow control di Golang
1. Validasi (if, switch)
2. Perulangan (hanya for)
*/
