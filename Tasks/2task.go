/* TODO: Konversi Suhu*/

package main

type DataCuaca struct {
	Suhu    float32
	Lokasi  string
	kondisi string
}

func (data *DataCuaca) KlasifikasiKondisi() string {
	if data.Suhu < 18 {
		return "Dingin"
	} else if data.Suhu >= 18 && data.Suhu <= 25 {
		return "Hangat"
	} else {
		return "Panas"
	}
}

func (data *DataCuaca) CelciusToFahrenheit() float32 {
	fahrenheit := (data.Suhu * 9 / 5) + 32
	return fahrenheit
}

func (data *DataCuaca) CelciusToReamur() float32 {
	reamur := data.Suhu * 4 / 5
	return reamur
}

func main() {
	dataCuaca := DataCuaca{
		Suhu:   25.0,
		Lokasi: "Jakarta",
	}

	kondisi := dataCuaca.KlasifikasiKondisi()
	fahrenheit := dataCuaca.CelciusToFahrenheit()
	reamur := dataCuaca.CelciusToReamur()

	println("Lokasi:", dataCuaca.Lokasi)
	println("Suhu dalam Celcius:", dataCuaca.Suhu)
	println("Kondisi:", kondisi)
	println("Suhu dalam Fahrenheit:", fahrenheit)
	println("Suhu dalam Reamur:", reamur)

}
