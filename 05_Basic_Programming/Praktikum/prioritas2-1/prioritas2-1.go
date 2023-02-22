package main

import "fmt"

func main() {
	input := 0

	fmt.Print("Input: ")
	fmt.Scanln(&input)

	fmt.Println("Output: ")

	if input != 0 {
		for i := 0; i < input; i++ {
			for j := 0; j <= i; j++ {
				if j == 0 {
					for k := 0; k < input-i; k++ {
						fmt.Printf(" ")
					}
				}
				fmt.Printf(" *")
			}
			fmt.Print("\n")
		}
	} else if input < 0 || input == 0 {
		fmt.Println("Input tidak boleh 0")
	}
}
