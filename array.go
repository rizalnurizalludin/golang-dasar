package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Hera"
	names[1] = "Ainu"
	names[2] = "Tazkia"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names)
	fmt.Println(len(names))

	var values = [3]int{
		90,
		80,
		70,
	}

	fmt.Println(values)
	fmt.Println(len(values))

	var lagi = [90]int{}
	fmt.Println(len(lagi))
}
