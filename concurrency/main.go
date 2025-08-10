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

	channelGoRoutine <- true // mengirim signal bahwa goroutine telah selesai dengan nilai true
}

func main() {
	fmt.Println("Memulai pesanan...")

	// inisialisasi channel
	channel := make(chan bool, 2) // alokasi memory (chan) dan tipe datanya (string, bool atu dll)

	// goROutine Utama
	go cetakPesan("Nasi Goreng", channel)

	//goROutine kedua
	go cetakPesan("Mie Goreng", channel)

	// Menunggu goROutine selesai
	// Dengan cara menunggu beberapa milidetik
	time.Sleep(5 * time.Millisecond)
	fmt.Println("Selesai")
}

/*
 Pesanan 1 -> Waktu proses
 Pesanan 2 -> Waktu proses
 Program pesanan 1 dan 2 berjalan secara bersamaan
*/
