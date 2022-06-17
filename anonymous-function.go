package main

import "fmt"

type Blacklist func(string) bool

func main() {

	//anonymous function, function tanpa nama
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("Rizal", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})

}

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

/*func blacklistAdmin(name string) bool {
	return name == "admin"
}*/
