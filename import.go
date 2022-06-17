package main

import (
	"fmt"
	"golang-dasar/bantuan"
	"golang-dasar/calculator"
)

func main() {
	helllo := bantuan.Katakanhelo("rizal")
	fmt.Println(helllo, bantuan.Application)

	/**
	-error program not exported by package bantuan
	hello := bantuan.katakanhelo("rizal")
	fmt.Println(hello, bantuan.version)
	*/

	c := calculator.Nilai{40, 30, 10}
	fmt.Println(c.Perkalian())
}
