package main

import (
	"errors"
)

// TODO: function to divide two numbers with error handling
func divide(a, b float32) (float32, error) {
	// Base case
	if b == 0 {
		return 0, errors.New("cannot divide by zero") // error handling errors.New
	}

	return a / b, nil
}

func main() {

}

/* TODO: Notes
- Di Golang kita harus mengecek error kembalian dari awal sesuai dengan error handling yang telah kita buat
*/
