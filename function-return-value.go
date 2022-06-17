package main

import "fmt"

func main() {

	result := getHello("Rizal")
	fmt.Println(result)

	fmt.Println(getHello(""))

}

func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}
}
