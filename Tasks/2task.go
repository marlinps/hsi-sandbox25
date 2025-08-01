/* TODO: Konversi Suhu*/

package main

type DataCuaca struct {
	Suhu    float32
	Lokasi  string
	kondisi string
}

func klasifikasiKondisi(data *DataCuaca) string {
	if data.Suhu < 18 {
		return "Dingin"
	} else if data.Suhu >= 18 && data.Suhu <= 25 {
		return "Hangat"
	} else {
		return "Panas"
	}
}

func celciusToFahrenheitt(data *DataCuaca) float32 {
	fahrenheit := (data.Suhu * 9 / 5) + 32
	return fahrenheit
}

func celciusToReamurr(data *DataCuaca) float32 {
	reamur := data.Suhu * 4 / 5
	return reamur
}

func main() {

}
