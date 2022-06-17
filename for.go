package main

import "fmt"

func main() {

	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	//init statement;post statement
	for counters := 1; counters <= 10; counters++ {
		fmt.Println("Perulangan Ke", counters)
	}

	//for array
	slice := []string{"Rizal", "Nurizalludin", "Ganteng"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//penulisan for range array dan slice
	//for range, for untuk melakukan perulangan terhadap semua data collection
	/*i adalah index slice dan array
	i bisa diganti _(underscore)
	untuk menginisialisasikan index yang tidak akan dipakai*/
	for i, value := range slice {
		fmt.Println(i, value)
		//fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Rizal"
	person["address"] = "Sukabumi"
	person["region"] = "Islam"

	//penulisan for range untuk map
	for key, value := range person {
		fmt.Println(key, ":", value)
		fmt.Println(person)
	}

}
