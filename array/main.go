package main

import "fmt"

func main() {
	// Declare an array of integers with a fixed size
	var arrayAngka [10]int // index 0 to 9

	// Assign values to the array
	arrayAngka[0] = 1
	arrayAngka[1] = 2
	arrayAngka[2] = 3
	arrayAngka[3] = 4
	arrayAngka[4] = 5

	// Print the values of the array
	fmt.Println("Data index ke-0:", arrayAngka[0]) // Output: 1

	// Initialize an array of integers
	arrayAngka1 := [5]int{10, 20, 30, 40, 50}

	// Print the values of the initialized array
	fmt.Println("Data index ke-0:", arrayAngka1[0]) // Output: 10

	// initialize array with shorthand syntax
	arrayAngka2 := [...]int{100, 200, 300, 400, 500}

	// Print the values of the shorthand initialized array
	fmt.Println("Panjang arrayAngka2:", len(arrayAngka2)) // Output: 5
	fmt.Println("Data index ke-0:", arrayAngka2[0])       // Output: 100
}

// TODO: Array adalahh struktur data yang memiliki ukuran tetap dan tipe data yang sama.
// Array dapat diakses menggunakan indeks, dimulai dari 0.
// Ukuran array ditentukan saat deklarasi dan tidak dapat diubah setelahnya.
