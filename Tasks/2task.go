/* TODO: Konversi Suhu*/

package main

type DataCuaca struct {
	Suhu    float32
	Lokasi  string
	kondisi string
}

func (data *DataCuaca) klasifikasiKondisi() string {
	if data.Suhu < 18 {
		return "Dingin"
	} else if data.Suhu >= 18 && data.Suhu <= 25 {
		return "Hangat"
	} else {
		return "Panas"
	}
}

func (data *DataCuaca) celciusToFahrenheitt(data *DataCuaca) float32 {
	fahrenheit := (data.Suhu * 9 / 5) + 32
	return fahrenheit
}

func (data *DataCuaca) celciusToReamurr() float32 {
	reamur := data.Suhu * 4 / 5
	return reamur
}

func main() {
	dataCuaca := DataCuaca{
		Suhu:    30.0,
		Lokasi:  "Jakarta",
		kondisi: "Panas",
	}

	kondisi := dataCuaca.klasifikasiKondisi()
	fahrenheit := dataCuaca.celciusToFahrenheitt(&dataCuaca)
	reamur := dataCuaca.celciusToReamurr()

	println("Lokasi:", dataCuaca.Lokasi)
	println("Suhu dalam Celcius:", dataCuaca.Suhu)
	println("Kondisi:", kondisi)
	println("Suhu dalam Fahrenheit:", fahrenheit)
	println("Suhu dalam Reamur:", reamur)

}
