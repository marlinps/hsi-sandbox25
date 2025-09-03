package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
	fmt.Println("Bismillah, Nama Saya Marlin Purnama Sari")

	var name *string
	name = new(string)
	fmt.Println("Pointer Name:", name)
	*name = "Marlin Purnama Sari"
	fmt.Println("Nama:", *name)
}
