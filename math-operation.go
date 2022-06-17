package main

import (
	"fmt"
	"math"
)

func main() {

	var result = 8 * 4
	fmt.Println(result)

	var a float64 = 3
	var b float64 = 3
	var c float64 = a * b

	akar := math.Sqrt(c)

	fmt.Println(c)
	fmt.Println(akar)

	var x = 89
	x += 98
	fmt.Println(x)
	x++
	fmt.Println(x)
}
