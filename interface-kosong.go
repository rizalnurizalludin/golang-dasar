package main

import "fmt"

func main() {

	var data interface{} = Ups(2)
	fmt.Println(data)

}

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}
