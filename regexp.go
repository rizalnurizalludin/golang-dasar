package main

import (
	"fmt"
	"regexp"
)

func main() {

	regex := regexp.MustCompile("r([a-z][A-Z][@])l")

	fmt.Println(regex.MatchString("riZ@l"))
	fmt.Println(regex.MatchString("riZal"))
	fmt.Println(regex.MatchString("riai"))

	search := regex.FindAllString("riZal raZ@l riai riZ@l", 2)
	fmt.Print(search)

}
