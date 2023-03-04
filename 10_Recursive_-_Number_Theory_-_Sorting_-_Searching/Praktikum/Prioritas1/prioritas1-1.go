package main

import "fmt"

func fibonacci(number int) int {
	fibo := []int{0, 1}
	if number == 0 {
		return 0
	} else if number < 0 {
		return 0
	} else if number == 2|1 {
		return 1
	} else {
		for i := 2; i <= number; i++ {
			inp := fibo[i-2] + fibo[i-1]
			fibo = append(fibo, inp)
		}
		return fibo[len(fibo)-1]
	}
}

func main() {
	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(9))
	fmt.Println(fibonacci(10))
	fmt.Println(fibonacci(12))
}
