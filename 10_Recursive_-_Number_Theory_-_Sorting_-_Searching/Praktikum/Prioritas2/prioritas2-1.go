package main

import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {

	for i := 0; i < len(cards); i++ {
		for j := 0; j < 2; j++ {
			if cards[i][j] == deck[0] || cards[i][j] == deck[1] {
				return cards[i]
			}
		}
	}
	return "tutup kartu"
}

func main() {
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1}))
}
