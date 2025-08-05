package main

import "fmt"

// TODO: Mendefinisikan interface/kontrak berupa method-method yang harus dimiliki oleh Struct
// TODO: Konsep Interface adalah Polimorpisme, dimana kita dapat menggunakan berbagai jenis struct yang berbeda selama mereka mengimplementasikan nama method dan paramater yang sama dengan logika yang berbeda
// TODO: Memungkinkan untuk membuat method-method yang berstruktur sehingga memeungkinkan kita untuk membuat kode yang lebih fleksibel dan reusable

type FormPendaftaranInterface interface {
	ValidasiUsia(usia int) bool
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

// Implementasi untuk struct FormPendaftaran reguler
func (f FormPendaftaran) ValidasiUsia(usia int) bool { // harus sama dengan interface (nama dan parameter)
	if usia < 15 || usia > 75 {
		return false
	}

	return true
}

// Implementasi untuk struct FormPendaftaran usia senja
func (f FormPendafataranUsiaSenja) ValidasiUsia(usia int) bool {
	if usia > 75 {
		return true
	}

	return false
}

// function
func ValidasiUsiaForm(fInt FormPendaftaranInterface, usia int) bool {
	return fInt.ValidasiUsia(usia)
}

func main() {
	pendaftaran1 := FormPendaftaran{
		NamaLengkap: "Budi Santoso",
		Email:       "email@email.com",
		Usia:        80,
	}

	pendaftaran2 := FormPendafataranUsiaSenja{
		NamaLengkap:    "Siti Aminah",
		Email:          "sitiaminah@email.com",
		Usia:           80,
		PenyakitKronis: "Diabetes",
	}

	apakahUsiaValid := ValidasiUsiaForm(pendaftaran1, pendaftaran1.Usia)
	fmt.Println("Apakah usia valid ?", apakahUsiaValid)

	apakahUsiaValid2 := ValidasiUsiaForm(pendaftaran2, pendaftaran2.Usia)
	fmt.Println("Apakah usia Senja valid ?", apakahUsiaValid2)

}
