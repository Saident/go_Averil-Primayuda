package main

import (
	"fmt"
	"math"
)

func primeX (number int) int {
	arr := []int{}
	start := 2
    for len(arr) < number {
        if isPrime(start) {
            arr = append(arr, start)
        }
        start++
    }
	return arr[len(arr) - 1]
}

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	div := int(math.Sqrt(float64(number)))
	for i := 2; i <= div; i++ {
		if number % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}
