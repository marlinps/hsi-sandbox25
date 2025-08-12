package main

import "fmt"

var hitung int // TODO: sumber konflik data (data race situation)

func tambah() {
	hitung++
}

func main() {
	for i := 0; i < 100; i++ {
		go tambah()
	}

	fmt.Print("Hasil akhir: ", hitung)
}

/* TODO: OUPUT
output : 97
seharusnya 100, inilah dinamakan data-race dimana setiap Go Routine melakukan proses timpa menimpa di variabel yang sama

TODO: hasilnya terkadang tidak konsisten tergantung CPU dilaptop masing-masing
TODO: kenapa bisa demikian ? karena go routine berjalan secara bersamaan
TODO: solusinya dengan TODO: go run --race <namafile.go>
*/
