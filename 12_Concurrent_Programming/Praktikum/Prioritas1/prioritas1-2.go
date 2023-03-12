package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	//Ubah x disini
	x := 5
	go kelipatanTiga(ch, x)
	for v := range ch {
		fmt.Println(v)
	}
}

func kelipatanTiga(ch chan int, x int) {
	for i := 1; i <= x; i++ {
		ch <- i * 3
		time.Sleep(1 * time.Second)
	}
	close(ch)
}
