package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func main() {

	users := []User{
		{"Rizal", 25},
		{"Hiqni", 21},
		{"Ebon", 18},
		{"Hera", 20},
	}

	sort.Sort(UserSlice(users))

	fmt.Print(users)

}

func (usrSlice UserSlice) Len() int {
	return len(usrSlice)
}

func (usrSlice UserSlice) Less(i, j int) bool {
	return usrSlice[i].Age < usrSlice[j].Age
}

func (usrSlice UserSlice) Swap(i, j int) {
	usrSlice[i], usrSlice[j] = usrSlice[j], usrSlice[i]
}
