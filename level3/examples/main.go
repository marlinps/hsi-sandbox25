package main

type PesertaSandbox struct {
	NamaLengkap  string
	Gender       string
	Usia         int
	SudahMenikah bool
	Domisili     string
}

func main() {
	// Inisialisasi struct dengan zero value
	var pesertaStructVariable PesertaSandbox
	pesertaStructVariable.NamaLengkap = "Marlin Purnama Sari"
	pesertaStructVariable.Gender = "Perempuan"
	pesertaStructVariable.Usia = 33
	pesertaStructVariable.SudahMenikah = false
	pesertaStructVariable.Domisili = "Jakarta"

	pesertaStructVariable2 := PesertaSandbox{
		NamaLengkap:  "Ria Rahmawati",
		Gender:       "Perempuan",
		Usia:         25,
		SudahMenikah: false,
		Domisili:     "Palembang",
	}

}
