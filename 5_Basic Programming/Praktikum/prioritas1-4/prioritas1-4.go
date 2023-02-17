package main

import "fmt"

func main() {
	fmt.Println("Angka 1 - 100: ")

	for i := 1; i < 101; i++ {
		if i%3 == 0 {
			fmt.Print(" Fizz")
		} else if i%5 == 0 {
			fmt.Print(" Buzz")
		} else {
			fmt.Print(" ", i, " ")
		}
	}
}
