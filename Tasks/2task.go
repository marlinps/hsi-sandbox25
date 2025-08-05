/* TODO: Konversi Suhu*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func (d *DataCuaca) CelciusToFahrenheit() float32 {
	fahrenheit := (d.Suhu * 9 / 5) + 32
	return fahrenheit
}

func (d *DataCuaca) CelciusToReamur() float32 {
	reamur := d.Suhu * 4 / 5
	return reamur
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--- Konverter Suhu ---")

	fmt.Print("Masukkan lokasi pengukuran suhu: ")
	lokasi, _ := reader.ReadString('\n')
	lokasi = strings.TrimSpace(lokasi)

	fmt.Print("Masukkan suhu dalam Celcius: ")
	suhuInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Gagal membaca input suhu:", err)
		return
	}
	suhuInput = strings.TrimSpace(suhuInput)
	suhu, err := strconv.ParseFloat(suhuInput, 32)
	if err != nil {
		fmt.Println("Input tidak valid, hanya menerima angka:", err)
		return
	}

	dataCuaca := DataCuaca{
		Suhu:   float32(suhu),
		Lokasi: lokasi,
	}

	valid, err := dataCuaca.ValidasiLokasi()
	if err != nil || !valid {
		fmt.Println("Error:", err)
		return
	}

	kondisi := dataCuaca.KlasifikasiKondisi()
	fahrenheit := dataCuaca.CelciusToFahrenheit()
	reamur := dataCuaca.CelciusToReamur()

	fmt.Printf("Kondisi: %s\n", kondisi)
	fmt.Printf("Suhu dalam Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Suhu dalam Reamur    : %.2f\n", reamur)
}
