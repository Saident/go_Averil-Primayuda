package main

import "fmt"

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20})) //30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) //40
}

func Frog(jumps []int) int {
    n := len(jumps)
    cost := make([]int, n)
    cost[0], cost[1] = 0, neg(jumps[1]-jumps[0])
    for i := 2; i < n; i++ {
        cost[i] = min(
			cost[i-1]+neg(jumps[i]-jumps[i-1]), 
			cost[i-2]+neg(jumps[i]-jumps[i-2]),
		)
    }
    return cost[n-1]
}

func neg(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

