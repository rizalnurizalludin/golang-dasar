package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var ktpRizal NoKTP = "34111124556"
	var marriedStatus Married = false

	fmt.Println(ktpRizal)
	fmt.Println(marriedStatus)
}
