package main

import (
	"fmt"
	"time"
)

func cetakPesan(pesanan string) {
	for i := 0; i < 5; i++ {
		fmt.Println("Pesanan ke-", i+1, ":", pesanan)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Memulai pesanan...")

	// goROutine Utama
	go cetakPesan("Nasi Goreng")

	//goROutine kedua
	go cetakPesan("Mie Goreng")

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
