package main

import "fmt"

func main() {
	// TODO: Metode 1 menginisialisasi array dengan ukuran tetap
	var arrayAngka [10]int // index 0 to  n-1

	// Assign values to the array
	arrayAngka[0] = 1
	arrayAngka[1] = 2
	arrayAngka[2] = 3
	arrayAngka[3] = 4
	arrayAngka[4] = 5

	// Print the values of the array
	fmt.Println("Data index ke-0:", arrayAngka[0]) // Output: 1 dan seterusnya

	// TODO: Metode 2 menginisialisasi array dengan nilai awal
	arrayAngka1 := [5]int{10, 20, 30, 40, 50}

	// Print the values of the initialized array
	fmt.Println("Data index ke-2:", arrayAngka1[2]) // Output: 10

	// TODO: Cara lain untuk menginisialisasi array, ukuran otomatis dengan jumlah elemen yang diberikan
	arrayAngka2 := [...]int{100, 200, 300, 400, 500}

	// Print the values of the shorthand initialized array
	fmt.Println("Panjang arrayAngka2:", len(arrayAngka2)) // Output: 5
	fmt.Println("Data index ke-3:", arrayAngka2[3])       // Output: 100
}

/* TODO: Pondasi Pemahaman Array
	Penjelasan:
	1. Array adalah struktur data yang memiliki ukuran tetap dan tipe data yang sama.
	2. Array dapat diakses menggunakan indeks, dimulai dari 0.
 	3. Ukuran array ditentukan saat deklarasi dan tidak dapat diubah setelahnya.
*/
