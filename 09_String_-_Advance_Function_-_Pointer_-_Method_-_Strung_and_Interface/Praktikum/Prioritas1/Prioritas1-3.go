package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	if len(a) <= len(b) {
		if strings.Contains(b, a) {
			return a
		}else {
			return "String tidak cocok"
		}
	}else if len(b) <= len(a) {
		if strings.Contains(a, b) {
			return b
		}else{
			return "String tidak cocok"
		}
	}else{
		return "String tidak cocok"
	}
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
