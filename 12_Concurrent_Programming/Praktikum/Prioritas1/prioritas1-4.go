package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{29, 31, 11, 15, 8}
	var sum int
	var total int

	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(1)
	go add(arr, &wg, &mutex, &sum, &total)
	wg.Wait()

	fmt.Println("Panjang Array:", total, ",Total:", sum)
}

func add(arr []int, wg *sync.WaitGroup, mutex *sync.Mutex, sum *int, total *int) {
	defer wg.Done()
	mutex.Lock()
	for _, v := range arr {
		*sum += v
		*total += 1
	}
	mutex.Unlock()
}
