package main

import "fmt"

func main() {

	/*deklarasi panjang
	var person map[string]string = map[string]string

	[string]=key
	string=value
	*/

	//deklarasi singkat
	person := map[string]string{
		"nama":    "Rizal",
		"address": "Sukabumi",
	}

	//menambah data key baru
	person["tittle"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["address"])
	fmt.Println(person["tittle"])

	//deklarasi map tanpa data apapun
	var book = make(map[string]string)

	//input data map book
	book["title"] = "Belajar Go-lang"
	book["author"] = "Rizal"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	//menghapus data pada map
	//delete(map, key)
	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))

}
