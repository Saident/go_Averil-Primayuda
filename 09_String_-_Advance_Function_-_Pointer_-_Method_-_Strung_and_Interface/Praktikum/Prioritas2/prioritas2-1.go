package main

import (
	"fmt"
)

func caesar(offset int, input string) string {
	res := ""
	for _, v := range input {
		ascii := int(v)
		offsetAscii := ((ascii - 'a') + offset) % 26 + 'a'
		res += string(rune(offsetAscii))
	}

	return res
}


func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
  	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}