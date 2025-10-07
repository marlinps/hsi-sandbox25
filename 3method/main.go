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

// TODO: method yang bersifat pas by reference, sehingga dapat mengubah data asli, menggunakan pointer receiver (*FormPendaftaran)
func (f *FormPendaftaran) MerubahUsia(usiaBaru int) {
	f.Usia = usiaBaru
	fmt.Println("Usia telah diubah menjadi:", f.Usia)
}

// fungsi untuk mencetak nama lengkap dari FormPendaftaran
func MencetakNamaLengkap(f FormPendaftaran) {
	fmt.Println("Nama Lengkap:", f.NamaLengkap)
}

func MerubahUsia2(f FormPendaftaran, usiaBaru int) {
	f.Usia = usiaBaru
	fmt.Println("Usia telah diubah menjadi:", f.Usia)
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

	fmt.Println("Struct original:", pendaftaran1)

	// method pass by reference
	pendaftaran1.MerubahUsia(35)

	// fungsi untuk merubah usia
	// MerubahUsia2(pendaftaran1, 35)

	// value usia telah berubah
	fmt.Println("Usia setelah perubahan:", pendaftaran1.Usia)

	fmt.Println("Struct setelah perubahan:", pendaftaran1)
}
