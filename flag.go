package main

import (
	"flag"
	"fmt"
)

func main() {

	host := flag.String("host", "localhost", "Put your host name")
	user := flag.String("host", "localhost", "Put your user name")
	password := flag.String("host", "localhost", "Put your password name")

	flag.Parse()

	fmt.Println("Host : ", *host)
	fmt.Println("Username : ", *user)
	fmt.Println("Password : ", *password)

}
