/* TODO: Notes
main.go
Entry point dari aplikasi kita
Special Package dari Golang
*/

package main

import (
	"fmt"
	"praktikum-module/calculator"

	// TODO: go get rsc.io/quote
	// TODO: Import package dari luar (third-party package),package untuk mendapatkan quote lucu

	"praktikum-module/stringgenerator"
)

func main() {
	hasilTambah := calculator.TambahAngka(5, 3)
	fmt.Println("Hasil Penjumlahan:", hasilTambah)
	hasilKurang := calculator.KurangAngka(5, 3)
	fmt.Println("Hasil Pengurangan:", hasilKurang)

	pesanBaru := stringgenerator.GenerateGreeting("Sari")
	fmt.Println("Pesan Baru:", pesanBaru)

}
