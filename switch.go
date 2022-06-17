package main

import "fmt"

func main() {
	name := "Rizal"

	switch name {
	case "Rizal":
		fmt.Println("Hello Rizal")
	case "Hera":
		fmt.Println("Hello Hera")
	case "Mimiw":
		fmt.Println("Hello Mimiw")
		fmt.Println("Hello Alien")
	default:
		fmt.Println("Hello Pendatang Baru")
	}

	//switch dengan short statement
	switch panjang := len(name); panjang >= 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	//switch tanpa kondisi
	panjang := len(name)
	switch {
	case panjang > 10:
		fmt.Println("Nama Terlalu Panjang")
	case panjang > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}

}
