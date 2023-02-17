package main

import "fmt"

func main() {
	var input int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&input)

	if input%2 == 0 {
		fmt.Println("Angka input merupakan bilangan genap")
	} else {
		fmt.Println("Angka input merupakan bilangan ganjil")
	}
}
