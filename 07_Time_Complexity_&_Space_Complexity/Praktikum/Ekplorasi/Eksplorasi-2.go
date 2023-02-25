package main

import (
	"fmt"
	"math"
)

func pow(x, n int)int{
	var x1 float64 = float64(x)
	var n1 float64 = float64(n)

	res := math.Pow(x1, n1)
	return int(res)
}

func main(){
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
}