package main

import (
	"fmt"
)

func main() {
	var a = 70
	var b = 50
	x := tambahan(a, b)
	y := kurang(a, b)
	z := kali(tambahan(a, b), kurang(a, b))
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func tambahan(a int, b int) int {
	hasil := a + b
	return hasil
}

func kurang(a int, b int) int {
	hasil := a - b
	return hasil
}

func kali(a int, b int) int {
	hasil := a * b
	return hasil
}
