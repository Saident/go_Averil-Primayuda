package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64
}

func (car Car) PerkiraanJarak() float64 {
	var ratio float64 = 0

	switch car.Type {
	case "Pertamax Turbo":
		ratio = 2.2
	case "Pertamax":
		ratio = 2.0
	case "Pertalite":
		ratio = 1.5
	}

	jarakTempuh := car.FuelIn * float64(ratio)
	return jarakTempuh
}

func main() {
	myCar := Car{"Pertamax", 12.3}
	fmt.Println(myCar.PerkiraanJarak(), "mill")

	myCar2 := Car{"Pertalite", 8.5}
	fmt.Println(myCar2.PerkiraanJarak(), "mill")
}
