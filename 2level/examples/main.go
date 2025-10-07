package main

import (
	"errors"
	"fmt"
)

// TODO: function to divide two numbers with error handling
func divide(a, b float32) (float32, error) {
	// Base case
	if b == 0 {
		return 0, errors.New("cannot divide by zero") // error handling errors.New
	}
	return a / b, nil
}

func kalkulasiPembagian(pembilang, pembagi int) (int, int, error) {
	if pembagi == 0 {
		return 0, 0, errors.New("cannot divide by zero") // error
	}

	hasil := pembilang / pembagi
	sisa := pembilang % pembagi

	return hasil, sisa, nil
}

func main() {
	// Exampleof successfull division
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Pesan Error: ", err)
	} else {
		fmt.Println("Hasil pembagian 10/2 =", result)
	}

	// Example of division by zero
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Hasil pembagian 10/0 =", result)
	}

	// kembalian dengan 3 nilai
	hasil, sisa, err := kalkulasiPembagian(10, 2)
	if err != nil {
		fmt.Println("Pesan Error: ", err)
	} else {
		fmt.Println("Hasil pembagian 10/2 =", hasil, "dengan sisa =", sisa)
	}
}

/* TODO: Notes
- Di Golang kita harus mengecek error kembalian dari awal sesuai dengan error handling yang telah kita buat
*/
