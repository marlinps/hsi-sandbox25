/* TODO: Notes
main.go
Entry point dari aplikasi kita
Special Package dari Go
*/

package main

import (
	"fmt"
	"praktikum-module/calculator"
)

func main() {
	hasilTambah := calculator.TambahAngka(5, 3)
	fmt.Println("Hasil Penjumlahan:", hasilTambah)

	hasilKurang := calculator.KurangAngka(5, 3)
	fmt.Println("Hasil Pengurangan:", hasilKurang)
}
