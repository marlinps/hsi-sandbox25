package main

import "fmt"

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
}

/* TODO: standart penulisan di Golang menggunakan camelCase
TODO: Reminder dalam memilih tipe data jangan asal-asal harus dipertimbangkan sesuai kebutuhan, alokasi memory;server dsb
TODO: Flow control di Golang
1. Validasi (if, switch)
2. Perulangan (hanya for)
*/
