package main

import "fmt"

func toPascal(numRows int) [][]int {
	output := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		arr := make([]int, i+1)
		arr[0], arr[len(arr)-1] = 1, 1
		output[i] = arr
		for j := 0; j < len(arr); j++ {
			if i > 1 {
				if j == 0 || j == len(arr)-1 {
					arr[j] = 1
				} else {
					arr[j] = output[i-1][j-1] + output[i-1][j]
				}
			}
		}
		output[i] = arr
	}
	return output
}

func main() {
	fmt.Println(toPascal(1))
	fmt.Println(toPascal(2))
	fmt.Println(toPascal(3))
	fmt.Println(toPascal(5))
}
