package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	merged := append(arrayA, arrayB...)

	uniqValue := make(map[string]int)
	res := []string{}

	for _, v := range merged {
		if uniqValue[v] == 0 {
			uniqValue[v] += 1
			res = append(res, v)
		}

	}
	return res
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
