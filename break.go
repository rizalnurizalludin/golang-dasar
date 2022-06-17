package main

import "fmt"

func main() {

	//break digunaakan untuk menghentikan seluruh perulangan
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke ", i)
	}
}
