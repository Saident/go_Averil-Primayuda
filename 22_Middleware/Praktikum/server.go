package main

import (
	"praktikum_22/route"
)

func main() {
	e := route.New()
	e.Logger.Fatal(e.Start(":8000"))
}
