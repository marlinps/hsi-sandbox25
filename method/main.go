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
// method yang berisifat pass by value, sehingga tidak mengubah data asli
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

	pendaftaran1.MencetakNamaLengkap()
	MencetakNamaLengkap(pendaftaran1)

	//method pass by reference
	pendaftaran1.MerubahUsia(35)
	fmt.Println("Usia setelah diubah:", pendaftaran1.Usia) // Output: 35
}

//TODO: keuntungan suatu method pada struct
// Penjelasan:
// 1. Method pada struct memungkinkan kita untuk mengelompokkan fungsi yang berhubungan dengan data dalam struct tersebut.
// 2. Dengan menggunakan method, kita dapat mengakses dan memanipulasi data dari struct dengan cara yang lebih terstruktur.
// 3. Method juga membantu dalam menjaga enkapsulasi, sehingga data dalam struct tidak dapat diakses secara langsung dari luar struct.
