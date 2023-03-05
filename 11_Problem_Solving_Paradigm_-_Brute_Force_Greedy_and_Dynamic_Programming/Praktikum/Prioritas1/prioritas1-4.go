package main

import "fmt"

func simpleEquations(a, b, c int) {
	//Pake complete search
	numSearch := 100
	x, y, z := 0, 0, 0
	for i := 1; i <= numSearch; i++ {
		x = i
		for j := 1; j <= numSearch; j++ {
			y = j
			for k := 1; k <= numSearch; k++ {
				z = k
				if x+y+z == a {
					if x*y*z == b {
						if x*x+y*y+z*z == c {
							fmt.Println(x, y, z)
							return
						}
					}
				}
			}
		}
	}
	fmt.Println("no solution")
}

func main() {
	simpleEquations(1, 2, 3)  // no solution
	simpleEquations(6, 6, 14) // 1 2 3
}
