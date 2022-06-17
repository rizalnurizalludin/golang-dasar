package main

import (
	"fmt"
	"math"
)

func main() {

	//membulatkan nilai ke yang paling dekat
	fmt.Println(math.Round(1.7))
	fmt.Println(math.Round(1.3))

	// membulatkan nilai ke yang paling rendah
	fmt.Println(math.Floor(1.7))

	// membulatkan nilai ke yang paling tinggi
	fmt.Println(math.Ceil(1.3))

	fmt.Println(math.Max(10, 34))
	fmt.Println(math.Min(10, 34))

}
