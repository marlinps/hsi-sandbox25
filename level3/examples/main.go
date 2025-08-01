package main

type PesertaSandbox struct {
	NamaLengkap  string
	Gender       string
	Usia         int
	SudahMenikah bool
	Domisili     string
}

func main() {
	// TODO: Inisialisasi struct dengan zero value
	var pesertaStructVariable PesertaSandbox // Nilai awalnya akan diisi oleh zero value
	pesertaStructVariable.NamaLengkap = "Marlin Purnama Sari"
	pesertaStructVariable.Gender = "Perempuan"
	pesertaStructVariable.Usia = 33
	pesertaStructVariable.SudahMenikah = false
	pesertaStructVariable.Domisili = "Jakarta"

	// TODO: Inisialisasi struct dengan literal
	pesertaStructVariable2 := PesertaSandbox{
		NamaLengkap:  "Ria Rahmawati",
		Gender:       "Perempuan",
		Usia:         25,
		SudahMenikah: false,
		Domisili:     "Palembang",
	}

}
