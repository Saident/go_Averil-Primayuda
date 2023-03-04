package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int
}

type pairs []pair

func MostAppearItem(items []string) []pair {
	uniqValue := make(map[string]int)
	for _, v := range items {
		if _, ok := uniqValue[v]; ok {
			uniqValue[v]++
		} else {
			uniqValue[v] = 1
		}
	}

	pairs := make(pairs, len(uniqValue))
	i := 0
	for k, v := range uniqValue {
		pairs[i] = pair{name: k + " ->", count: v}
		i++
	}

	//SORT
	length := len(pairs)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if pairs[j].count > pairs[j+1].count {
				pairs[j], pairs[j+1] = pairs[j+1], pairs[j]
			}
		}
	}
	return pairs
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}
