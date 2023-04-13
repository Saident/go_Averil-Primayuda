package main

import (
	"praktikum_23_2/route"
)

func main() {
	e := route.New()
	e.Logger.Fatal(e.Start(":8000"))
}
