package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func celciusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func celciusToReamur(celsius float64) float64 {
	return celsius * 4 / 5
}

func validateInput(input string) (float64, error) {
	celsius, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("input tidak Valid, hanya menerima angka: %v", err)
	}
	return celsius, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("masukkan suhu dalam Celsius: ")

	celsiusStr, _ := reader.ReadString('\n')

	celsiusStr = strings.TrimSpace(celsiusStr)

	celsius, err := validateInput(celsiusStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fahrenheit := celciusToFahrenheit(celsius)
	reamur := celciusToReamur(celsius)

	fmt.Printf("Suhu dalam Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Suhu dalam Reamur: %.2f\n", reamur)
}
