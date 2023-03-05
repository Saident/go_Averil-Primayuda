package main

import (
	"fmt"
	"strconv"
)

func ans(number int) []string {
	ans := make([]string, number+1)
	for i := 0; i <= number; i++ {
		ans[i] = setBit(i)
	}
	return ans
}

func setBit(number int) string {
	out := ""
	if number == 0 {
		return "0"
	} else if number == 1 {
		return "1"
	} else {
		for number > 0 {
			for number > 0 {
				ext := number & 1
				out = strconv.Itoa(ext) + out
				number >>= 1
			}
		}
		return out
	}
}

func main() {
	fmt.Println(ans(2))
	fmt.Println(ans(3))
	fmt.Println(ans(4))
}
