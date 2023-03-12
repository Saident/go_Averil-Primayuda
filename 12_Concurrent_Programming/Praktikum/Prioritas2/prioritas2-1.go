package main

import (
	"fmt"
	"sync"
)

func main() {
	text := "Lorem Ipsum dolor sit amet"

	freq := make(map[string]int)

	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(1)
	go func(wg *sync.WaitGroup, mutex *sync.Mutex, text string) {
		defer wg.Done()
		mutex.Lock()
		for _, char := range text {
			freq[string(char)]++
		}
		for char, total := range freq {
			fmt.Println(char, ":", total)
		}
		mutex.Unlock()
	}(&wg, &mutex, text)
	wg.Wait()
}
