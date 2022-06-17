package main

import "fmt"

func main() {
	sayHelloTo("Rizal", "Nurizalludin")

	NamaDepan := "Rizal"
	sayHelloTo(NamaDepan, "Nurizalludin")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
