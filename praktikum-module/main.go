/* TODO: Notes
main.go
Entry point dari aplikasi kita
Special Package dari Go
*/

package main

import (
	"fmt"
	"praktikum-module/calculator"

	// TODO: go get rsc.io/quote
	"rsc.io/quote" // TODO: Import package dari luar (third-party package), package untuk mendapatkan quote lucu
)

func main() {
	hasilTambah := calculator.TambahAngka(5, 3)
	fmt.Println("Hasil Penjumlahan:", hasilTambah)
	hasilKurang := calculator.KurangAngka(5, 3)
	fmt.Println("Hasil Pengurangan:", hasilKurang)

	pesan := quote.Hello()
	fmt.Println("Pesan dari quote package:", pesan)
}
