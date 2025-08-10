/* TODO: Konversi Suhu*/

package main

import (
	"fmt"
	"strconv"
)

type DataCuaca struct {
	Suhu    float32
	Lokasi  string
	kondisi string
}

func (d *DataCuaca) ValidasiLokasi() (bool, error) {
	if d.Lokasi == "" {
		return false, fmt.Errorf("Lokasi tidak boleh kosong")
	}

	for _, c := range d.Lokasi {
		if c >= '0' && c <= '9' {
			return false, fmt.Errorf("Data lokasi tidak boleh ada angka: %s", d.Lokasi)
		}
	}

	return true, nil
}

func (d *DataCuaca) KlasifikasiKondisi() string {
	if d.Suhu < 18 {
		return "Dingin"
	} else if d.Suhu >= 18 && d.Suhu <= 25 {
		return "Hangat"
	} else {
		return "Panas"
	}
}

func (data *DataCuaca) celciusToFahrenheit() float32 {
	fahrenheit := (data.Suhu * 9 / 5) + 32
	return fahrenheit
}

func (data *DataCuaca) celciusToReamur() float32 {
	reamur := data.Suhu * 4 / 5
	return reamur
}

func main() {
	var suhuInput string
	println("Masukkan suhu dalam Celcius:")
	_, err := fmt.Scanln(&suhuInput)
	if err != nil {
		println("Terjadi kesalahan input.")
		return
	}

	suhu, err := strconv.ParseFloat(suhuInput, 32)
	if err != nil {
		println("Input suhu harus berupa angka.")
		return
	}

	dataCuaca := DataCuaca{
		Suhu:    (float32)(suhu),
		Lokasi:  "Jakarta",
		kondisi: "",
	}

	kondisi := dataCuaca.klasifikasiKondisi()
	fahrenheit := dataCuaca.celciusToFahrenheit()
	reamur := dataCuaca.celciusToReamur()

	println("Lokasi:", dataCuaca.Lokasi)
	println("Suhu dalam Celcius:", dataCuaca.Suhu)
	println("Kondisi:", kondisi)
	println("Suhu dalam Fahrenheit:", fahrenheit)
	println("Suhu dalam Reamur:", reamur)
}
