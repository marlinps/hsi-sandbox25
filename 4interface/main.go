package main

import "fmt"

// TODO: Mendefinisikan interface/kontrak berupa method-method yang harus dimiliki oleh Struct, semua struct harus mengimplementasikan method-method ini, kalau tidak maka akan terjadi error
// TODO: Konsep Interface adalah Polimorfism, dimana kita dapat menggunakan berbagai jenis struct yang berbeda selama mereka mengimplementasikan nama method dan paramater yang sama dengan logika yang berbeda
// TODO: Memungkinkan untuk membuat method-method yang berstruktur sehingga memeungkinkan kita untuk membuat kode yang lebih fleksibel dan reusable

type FormPendaftaranInterface interface {
	ValidasiUsia(usia int) bool
	ValidasiGender(gender string) bool
}

type FormPendaftaran struct {
	NamaLengkap string
	Email       string
	Gender      string
	Usia        int
}

type FormPendafataranUsiaSenja struct {
	NamaLengkap    string
	Email          string
	Gender         string
	Usia           int
	PenyakitKronis string
}

// Implementasi method untuk struct FormPendaftaran reguler
func (f FormPendaftaran) ValidasiUsia(usia int) bool { // harus sama dengan interface (nama dan parameter)
	if usia < 15 || usia > 75 {
		return false
	}

	return true
}

func (f FormPendaftaran) ValidasiGender(gender string) bool {
	if gender != "Laki-laki" && gender != "Perempuan" {
		return false
	}
	return true
}

// Implementasi method untuk struct FormPendaftaran usia senja
func (f FormPendafataranUsiaSenja) ValidasiUsia(usia int) bool {
	if usia > 75 {
		return true
	}
	return false
}

func (f FormPendafataranUsiaSenja) ValidasiGender(gender string) bool {
	if gender != "Laki-laki" && gender != "Perempuan" {
		return false
	}
	return true
}

// Implementasi fungsi untuk validasi usia dan gender
func ValidasiUsiaForm(fInt FormPendaftaranInterface, usia int) bool {
	return fInt.ValidasiUsia(usia)
}

func ValidasiGenderForm(fInt FormPendaftaranInterface, gender string) bool {
	return fInt.ValidasiGender(gender)
}

func main() {
	p1 := FormPendaftaran{
		NamaLengkap: "Budi Santoso",
		Email:       "email@email.com",
		Gender:      "Laki-laki",
		Usia:        80,
	}

	p2 := FormPendafataranUsiaSenja{
		NamaLengkap:    "Siti Aminah",
		Email:          "sitiaminah@email.com",
		Gender:         "Perempuan",
		Usia:           80,
		PenyakitKronis: "Diabetes",
	}

	apakahUsiaValid := ValidasiUsiaForm(p1, p1.Usia)
	fmt.Println("Apakah usia valid ?", apakahUsiaValid)

	apakahGenderValid := ValidasiGenderForm(p1, p1.Gender)
	fmt.Println("Apakah gender valid ?", apakahGenderValid)

	apakahUsiaValid2 := ValidasiUsiaForm(p2, p2.Usia)
	fmt.Println("Apakah usia Senja valid ?", apakahUsiaValid2)

	apakahGenderValid2 := ValidasiGenderForm(p2, p2.Gender)
	fmt.Println("Apakah gender Senja valid ?", apakahGenderValid2)

}
