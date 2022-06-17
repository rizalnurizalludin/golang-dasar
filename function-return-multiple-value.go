package main

import "fmt"

func main() {

	firstname, lastname := getFullname()
	fmt.Println(firstname)
	fmt.Println(lastname)

	//menghiraukan return value
	_, lastName := getFullname()
	fmt.Println(lastName)

}

func getFullname() (string, string) {
	return "Rizal", "Nurizalludin"
}
