package main

import "fmt"

type RawData struct {
	data    []int
	channel chan int
}

func hitungData(rd RawData) {
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

	total1 := RawData{data: angkaGenap, channel: c1}
	total2 := RawData{data: angkaGanjil, channel: c2}

	go hitungData(total1)
	go hitungData(total2)

	totalAngkaGenap := <-c1  // tunggu sampai goroutine pertama selesai, ambil data dari channel 1
	totalAngkaGanjil := <-c2 // tunggu sampai goroutine pertama selesai, ambil data dari channel 2

	fmt.Println("Total Angka Genap:", totalAngkaGenap)
	fmt.Println("Total Angka Ganjil:", totalAngkaGanjil)

	// Total keseluruhan
	fmt.Println("Total Keseluruhan:", totalAngkaGenap+totalAngkaGanjil) // harus menunggu kedua goroutine selesai sebelum mencetak total keseluruhan
}
