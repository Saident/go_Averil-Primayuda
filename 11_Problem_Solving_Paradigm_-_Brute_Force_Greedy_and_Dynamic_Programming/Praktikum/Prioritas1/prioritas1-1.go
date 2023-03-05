package main

import (
	"fmt"
	"strconv"
)

func ans(number int64) []string {
	ans := make([]string, number+1)
	var i int64
	for i = 0; i <= number; i++ {
		ans[i] = strconv.FormatInt(i, 2)
	}
	return ans
}

func main() {
	fmt.Println(ans(2))
	fmt.Println(ans(3))
	fmt.Println(ans(4))
}
