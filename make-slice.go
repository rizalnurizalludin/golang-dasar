package main

import "fmt"

func main() {
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Rizal"
	newSlice[1] = "Nurizalludin"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
	fmt.Println(len(copySlice))
	fmt.Println(cap(copySlice))

	//penulisan kode array harus masukan panjang kapasitas array
	iniArray := [5]int{1, 2, 3, 4, 5}

	//penulisan slice bisa dikosongkan dan diinisialisasikan ...
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
