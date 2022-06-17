package main

import "fmt"

func main() {
	name := "Rizal"

	if name == "Rizal" {
		fmt.Println("Hello Rizal")
	} else if name == "Hera" {
		fmt.Println("Hello Hera")
	} else {
		fmt.Println("Hello Pendatang Baru")
	}

	//if dengan short statement
	if panjang := len(name); panjang >= 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Benar")
	}
}
