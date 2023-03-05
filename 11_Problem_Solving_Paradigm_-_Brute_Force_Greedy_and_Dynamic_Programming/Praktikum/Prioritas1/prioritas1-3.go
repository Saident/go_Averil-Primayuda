package main

import "fmt"

func fibo(numRows int) int {
	arr := make([]int, numRows+1)
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			arr[i] = 0
		} else if i == 1 {
			arr[i] = 1
		} else if i == 2 {
			arr[i] = 1
		} else {
			arr[i] = arr[i-2] + arr[i-1]
		}
	}
	return arr[len(arr)-1]
}

func main() {
	fmt.Println(fibo(0))
	fmt.Println(fibo(2))
	fmt.Println(fibo(3))
	fmt.Println(fibo(5))
	fmt.Println(fibo(6))
	fmt.Println(fibo(7))
	fmt.Println(fibo(9))
}
