package main

import "fmt"

func main() {

	counter := 0

	/*closures adalah kemampuan sebuah function berinteraksi
	dengan data-data disekitarnya dalam scope yang sama
	*/
	increment := func() {
		fmt.Println("increment")
		counter++
	}

	/*counter dipengaruhi oleh function closures increment
	dan berubah seiring function digunakan*/
	increment()
	increment()
	fmt.Println(counter)
}
