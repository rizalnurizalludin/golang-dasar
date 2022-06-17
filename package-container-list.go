package main

import (
	"container/list"
	"fmt"
)

//double link list
func main() {

	data := list.New()
	data.PushBack("Rizal")
	data.PushBack("Nurizalludin")
	data.PushBack("Ganteng")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	// dari depan ke belakang
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	//dari belakang ke depan
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

}
