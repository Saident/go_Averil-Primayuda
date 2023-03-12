package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	//Ubah x disini
	x := 5
	go cetakKelipatan(x, &wg)
	wg.Wait()
}

func cetakKelipatan(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 1
	for {
		fmt.Println(i * x)
		i++
		time.Sleep(3 * time.Second)
	}
}


