package main

import "fmt"

func main() {
	var atas, bawah, tinggi float64

	fmt.Print("Masukkan panjang sisi atas trapesium: ")
	fmt.Scanln(&atas)

	fmt.Print("Masukkan panjang sisi bawah trapesium: ")
	fmt.Scanln(&bawah)

	fmt.Print("Masukkan panjang sisi tinggi trapesium: ")
	fmt.Scanln(&tinggi)

	luas := (atas + bawah) * tinggi / 2
	fmt.Println("Luas Trapesium adalah:", luas)
}
