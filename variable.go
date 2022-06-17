package main

import (
	"fmt"
)

func main() {
	var name string

	name = "Rizal Nurizalludin"
	fmt.Println(name)
	fmt.Println(len(name))

	var girlfriendName = "Hera Ainu Tazkia"
	fmt.Println(girlfriendName)

	var age = 20
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	country = "Inggris"
	fmt.Println(country)

	var (
		firsName = "Rizal"
		lastName = "Nurizalludin"
	)
	fmt.Println(firsName + " " + lastName)
}
