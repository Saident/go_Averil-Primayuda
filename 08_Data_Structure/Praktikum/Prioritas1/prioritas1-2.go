package main

import "fmt"

func mapping(slice []string) map[string]int {
	uniqValue := make(map[string]int)
	for _, v := range slice {
		if _, ok := uniqValue[v]; ok {
			uniqValue[v]++
		} else {
			uniqValue[v] = 1
		}
	}
	return uniqValue
}

func main() {
	fmt.Println(mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(mapping([]string{}))
}
