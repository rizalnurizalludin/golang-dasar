//merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
//sering digunakan ketika bertemu dengan interfaxce kosong
package main

import "fmt"

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	//type assertsion menggunakan switch
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	case bool:
		fmt.Println("Boolean", value)
	default:
		fmt.Println("Unknown", value)
	}
}

func random() interface{} {
	return 3.4
}
