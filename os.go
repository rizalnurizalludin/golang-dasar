package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Arguments: ", args)
	fmt.Println(args[1])
	fmt.Println(args[2])

	name, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(name)
	}

	username := os.Getenv("APP_USERNAME")
	pass := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(pass)
}
