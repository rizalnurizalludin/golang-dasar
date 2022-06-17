package main

import (
	"errors"
	"fmt"
)

func main() {

	//fmt.Println(Pembagian(10, 3))

	hasil, err := Pembagian(100, 10)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}

}

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembangian dengan Nol")
	} else {
		return nilai / pembagi, nil
	}
}
