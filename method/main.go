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

	//method pass by reference
	pendaftaran1.MerubahUsia(35)
	fmt.Println("Usia setelah diubah:", pendaftaran1.Usia) // Output: 35
}

/*TODO: keuntungan suatu method pada struct
Penjelasan:
1. Method pada struct memungkinkan kita untuk mengelompokkan fungsi yang berhubungan dengan data dalam struct tersebut.
2. Dengan menggunakan method, kita dapat mengakses dan memanipulasi data dari struct dengan cara yang lebih terstruktur.
3. Method juga membantu dalam menjaga enkapsulasi, sehingga data dalam struct tidak dapat diakses secara langsung dari luar struct.
4. Method dapat memiliki receiver yang memungkinkan kita untuk mengubah data dalam struct secara langsung, seperti pada contoh MerubahUsia.
5. Dengan menggunakan pointer receiver, kita dapat mengubah nilai asli dari struct yang dipanggil, sehingga perubahan tersebut akan terlihat di luar method.
6. Method juga dapat digunakan untuk membuat kode lebih modular dan mudah dibaca, karena kita dapat memberikan nama yang jelas pada setiap method sesuai dengan fungsinya. (Sehingga lebih mudah dipahami oleh pembaca kode).
7. Dengan menggunakan method, kita dapat menghindari duplikasi kode, karena kita dapat menggunakan kembali method yang sama pada beberapa instance dari struct yang berbeda.
8. Method juga memungkinkan kita untuk membuat antarmuka (interface) yang dapat digunakan untuk mengimplementasikan polimorfisme dalam program.
9. Dengan menggunakan method, kita dapat membuat kode yang lebih mudah untuk diuji dan dipelihara, karena setiap method dapat diuji secara terpisah.
10. Method pada struct juga memungkinkan kita untuk membuat kode yang lebih fleksibel dan dapat digunakan kembali di berbagai konteks.
*/
