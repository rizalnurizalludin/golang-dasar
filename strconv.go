package main

import (
	"fmt"
	"strconv"
)

// string convers
func main() {

	//konversi dari string ke tipe data lain
	boolean, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(boolean)
	}

	//binary =2, octal =8 ,desimal = 10, hexa desimal =16
	number, err := strconv.ParseInt("1000000", 10, 64)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(number)
	}

	//konversi dari tipe data lain ke string
	//binary =2, octal =8, decimal = 10, hexadecimal =16
	value := strconv.FormatInt(1000000, 2)
	fmt.Println(value)

	// konversi ke number lebih mudah
	valueInt, _ := strconv.Atoi("200000")
	fmt.Println(valueInt)

	//konversi dari number ke string
	valueStr := strconv.Itoa(30000)
	fmt.Println(valueStr)
}
