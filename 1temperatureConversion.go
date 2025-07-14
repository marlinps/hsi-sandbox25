package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func celciusToFahrenheit(celsius float32) float32 {
	return (celsius * 9 / 5) + 32
}

func celciusToReamur(celsius float32) float32 {
	return celsius * 4 / 5
}

func validateInput(input string) (float32, error) {
	celsius, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return 0, fmt.Errorf("input tidak Valid, hanya menerima angka: %v", err)
	}

	return float32(celsius), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--- Konverter Suhu ---")

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
