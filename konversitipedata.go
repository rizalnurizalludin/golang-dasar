package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilai64 = int64(nilai32)

	var nilai8 = int8(nilai32)

	fmt.Println(nilai8)
	fmt.Println(nilai32)
	fmt.Println(nilai64)

	var name = "Rizal"
	var e byte = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
