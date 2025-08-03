package main

import "fmt"

// inisialisasi struct untuk form pendaftaran
type FormPendaftaran struct {
	NamaLengkap string
	Email       string
	Gender      string
	Usia        int
}

// method untuk menampilkan informasi pendaftaran
// menggunakan receiver pada struct (f FormPendaftaran) biasanya ditulis satu karakter saja (f)
func (f FormPendaftaran) MencetakNamaLengkap() {
	fmt.Println("Nama Lengkap:", f.NamaLengkap)
}

func main() {

}
