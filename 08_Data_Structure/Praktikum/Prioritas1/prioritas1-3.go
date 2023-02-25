package main

import (
	"fmt"
	"strings"
	"strconv"
)

func munculSekali(angka string) []int {
	res := []int{}
	arrString := strings.Split(angka, "")
	uniqValue := make(map[int]int)

    arrInt := make([]int, len(arrString))
    for i, v := range arrString {
        conv, err := strconv.Atoi(v)
        if err != nil {
            panic(err)
        }
        arrInt[i] = conv
    }

	for _, v := range arrInt {
		if _, ok := uniqValue[v]; ok {
			uniqValue[v]++
		} else {
			uniqValue[v] = 1
		}
	}

	for _, v := range arrInt {
		if uniqValue[v] == 1 {
			res = append(res, v)
		}
	}

	
	return res
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
