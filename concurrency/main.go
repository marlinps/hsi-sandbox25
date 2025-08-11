package main

import (
	"fmt"
	"time"
)

func cetakPesan(pesanan string, channelGoRoutine chan bool) {
	for i := 0; i < 5; i++ {
		fmt.Println("Pesanan ke-", i+1, ":", pesanan)
		time.Sleep(1 * time.Millisecond)
	}

	channelGoRoutine <- true // mengirim data ke channel bahwa goroutine telah selesai dengan nilai true
}

func main() {
	fmt.Println("Memulai pesanan...")

	// inisialisasi channel
	channel := make(chan bool, 2) // alokasi memory make(chan) dan tipe datanya (string, bool atau dll)

	// goROutine Utama
	go cetakPesan("Nasi Goreng", channel)

	// goROutine kedua
	go cetakPesan("Telur Mata Sapi", channel)

	// Menunggu goROutine selesai
	// Dengan cara menunggu beberapa milidetik
	time.Sleep(5 * time.Millisecond)
	// Tidak akan mencetak pesan "Selesai" sebelum channel menerima data dan bernilai true
	fmt.Println("Selesai")
}

/*
 Pesanan 1 -> Waktu proses
 Pesanan 2 -> Waktu proses
 TODO: Program pesanan 1 dan 2 berjalan secara bersamaan tanpa menunggu satu sama lain. (Concurrency)
*/

/*
TODO: Output:
Memulai pesanan...
Pesanan ke- 1 : Nasi Goreng
Pesanan ke- 2 : Nasi Goreng
Pesanan ke- 3 : Nasi Goreng
Pesanan ke- 4 : Nasi Goreng
Pesanan ke- 5 : Nasi Goreng
Pesanan ke- 1 : Telur Mata Sapi
Pesanan ke- 2 : Telur Mata Sapi

TODO: Penjelasan output
meskipun output seolah-olah dicetak berurutan, namun sebenarnya kedua goroutine berjalan secara bersamaan.
Pesanan "Nasi Goreng" dan "Telur Mata Sapi" dicetak secara bersamaan, dengan masing-masing goroutine mencetak 5 pesan.
*/
