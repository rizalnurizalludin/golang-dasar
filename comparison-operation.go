package main

import "fmt"

func main() {
	a := "Rizal"
	b := "Nurizalludin"
	c := "Rizal"

	x := 100
	y := 200

	result1 := a == b
	result2 := a == c
	result3 := x > y
	result4 := x < y

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x == y)
	fmt.Println(x != y)
}
