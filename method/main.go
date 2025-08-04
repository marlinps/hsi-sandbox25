package main

import "fmt"

// inisialisasi struct bernama form pendaftaran
type FormPendaftaran struct {
	NamaLengkap string
	Email       string
	Gender      string
	Usia        int
}

// TODO: method untuk menampilkan informasi pendaftaran
// TODO: menggunakan receiver pada struct (f FormPendaftaran) biasanya ditulis satu karakter saja (f)
// TODO: method yang bersifat pass by value, sehingga tidak mengubah data asli
func (f FormPendaftaran) MencetakNamaLengkap() {
	fmt.Println("Nama Lengkap:", f.NamaLengkap)
}

// method yang bersifat pas by reference, sehingga dapat mengubah data asli
func (f *FormPendaftaran) MerubahUsia(usiaBaru int) {
	f.Usia = usiaBaru
}

// fungsi untuk mencetak nama lengkap dari FormPendaftaran
func MencetakNamaLengkap(f FormPendaftaran) {
	fmt.Println("Nama Lengkap:", f.NamaLengkap)
}

func main() {
	pendaftaran1 := FormPendaftaran{
		NamaLengkap: "John Doe",
		Email:       "john.doe@example.com",
		Gender:      "Laki-laki",
		Usia:        30,
	}

	// method untuk mencetak nama lengkap
	pendaftaran1.MencetakNamaLengkap()

	// fungsi untuk mencetak nama lengkap
	MencetakNamaLengkap(pendaftaran1)

	// method pass by reference
	pendaftaran1.MerubahUsia(35)
	fmt.Println("Usia setelah diubah:", pendaftaran1.Usia) // Output: 35
}
