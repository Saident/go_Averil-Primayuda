package main

import (
	"praktikum_21/route"
)

func main() {
	e := route.New()
	e.Logger.Fatal(e.Start(":8000"))
}
