package main

import (
	"fmt"
)

func main() {
	var (
		myfirstname = "Rizal"
		mylastname  = "Nurizalludin"
	)
	fullname := myfirstname + " " + mylastname

	fmt.Println(myfirstname)
	fmt.Println(len(myfirstname))
	fmt.Println(fullname)
	fmt.Println(len(fullname))
	fmt.Println(fullname[:10])
	fmt.Println("Rizal Nurizalludin Manual")
}
