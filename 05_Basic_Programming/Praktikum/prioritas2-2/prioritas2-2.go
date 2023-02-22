package main

import "fmt"

func main() {
	input := 0

	fmt.Print("Angka: ")
	fmt.Scan(&input)

	for i := 1; i <= input; i++ {
		if input%i == 0 {
			fmt.Println(i)
		}
	}
}
