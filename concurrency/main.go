package main

import (
	"errors"
	"fmt"
)
type RawData struct {
	data         []int
	channel      chan int
	channelError chan error
}

func hitungData(rd RawData) {
	if len(rd.data) == 0 {
		rd.channel <- 0
		rd.channelError <- errors.New("Deretnya tidak boleh kosong")
		return
	}
	total := 0
	for _, value := range rd.data {
		total += value
	}

	rd.channel <- total // Send the total to the channel
}

func main() {
	angkaGenap := []int{2, 4, 6, 8, 10}
	angkaGanjil := []int{1, 3, 5, 7, 9}

	c1 := make(chan int)
	c2 := make(chan int)

	c1Error := make(chan error)
	c2Error := make(chan error)

	total1 := RawData{data: angkaGenap, channel: c1}  // mungkin berjalan 5 ms
	total2 := RawData{data: angkaGanjil, channel: c2} // mungkin berjalan 12 ms

	go hitungData(total1)
	go hitungData(total2)

	totalAngkaGenap := <-c1  // tunggu sampai goroutine pertama selesai, ambil data dari channel 1
	totalAngkaGanjil := <-c2 // tunggu sampai goroutine pertama selesai, ambil data dari channel 2

	fmt.Println("Total Angka Genap:", totalAngkaGenap)
	fmt.Println("Total Angka Ganjil:", totalAngkaGanjil)

	// Total keseluruhan
	fmt.Println("Total Keseluruhan:", totalAngkaGenap+totalAngkaGanjil) // harus menunggu kedua goroutine selesai sebelum mencetak total keseluruhan
}
