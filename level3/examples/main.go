package main

import "fmt"

type PesertaSandbox struct {
	NamaLengkap  string
	Gender       string
	Usia         int
	SudahMenikah bool
	Domisili     string
}

// TODO: struct bisa digunakan sebagai argumen di function
// TODO: akan diteruskan sbg pass by value, mengirimkan dengan nilai sehingga struct dibuat baru atau salin (copy)
func bertambahUsia(peserta PesertaSandbox) {
	peserta.Usia += 1
	fmt.Println("Usian peserta setelah melewati function bertambahUsia =", peserta.Usia)
}

// TODO: diteruskan sbg pass by reference
// TODO: struct tidak akan dibuat baru hanya mengirimkan alamat struct
// TODO: sehingga perubahan pada struct akan berpengaruh pada struct aslinya
func bertambahUsiaPointer(peserta *PesertaSandbox) {
	peserta.Usia += 1
	fmt.Println("Usian peserta setelah melewati function bertambahUsia =", peserta.Usia)
}

func main() {
	// TODO: Inisialisasi struct dengan zero value
	var pesertaStructVariable PesertaSandbox // Nilai awalnya akan diisi oleh zero value
	pesertaStructVariable.NamaLengkap = "Marlin Purnama Sari"
	pesertaStructVariable.Gender = "Perempuan"
	pesertaStructVariable.Usia = 33
	pesertaStructVariable.SudahMenikah = false
	pesertaStructVariable.Domisili = "Jakarta"

	// TODO: ketika kamu punya pointer ke struct, bisa langsung akses field dengan . — compiler otomatis dereferensi, jadi gak perlu mengubah nilai di struct aslinya
	pesertaStructVariablePointer := &pesertaStructVariable
	pesertaStructVariablePointer.NamaLengkap = "Marlin P.S"
	pesertaStructVariablePointer.Usia = 30

	fmt.Println(pesertaStructVariable.NamaLengkap)
	fmt.Println(pesertaStructVariable.Usia)

	// TODO: Inisialisasi struct dengan literal
	// TODO: Umumnya digunakan untuk mendapatkan Default Value dari  sebuah Struct (runtimenya lebih cepat dibandingkan inisialisasi diatas

	// pesertaStructVariable2 := PesertaSandbox{
	// 	NamaLengkap:  "Ria Rahmawati",
	// 	Gender:       "Perempuan",
	// 	Usia:         25,
	// 	SudahMenikah: false,
	// 	Domisili:     "Palembang",
	// }

	// Pass by Value, dimana Struct akan dicopy paste
	bertambahUsia(pesertaStructVariable) //34

	bertambahUsiaPointer(&pesertaStructVariable)

	// Struct aslinya tidak berubah
	fmt.Println(pesertaStructVariable.Usia) //33

	// TODO: Pointer
	// var alamatAsli string = "Duren Sawit Jakarta Timur"
	// var pointerAlamat *string = &alamatAsli // TODO: Pointer dengan tipe string menunjuk ke memory alamat asli

	// fmt.Println("Nilai dari variable alamatAsli :", alamatAsli)
	// fmt.Println("Nilai dari variable Pointer(Memory) alamat asli :", pointerAlamat)
	// fmt.Println("Isi dari variable pointerAlamat :", *pointerAlamat) // TODO: Akan mengembalikan nilai aslinya

}
