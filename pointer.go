package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	address1 := Address{
		City:     "Bandung",
		Province: "West Java",
		Country:  "Cimahi",
	}
	address2 := &address1
	address3 := &address1
	address4 := &address1
	address5 := new(Address)
	address6 := Address{
		City:     "Manchester City",
		Province: "Premiere League",
		Country:  "Indonesia",
	}
	//merubah field city address1
	address2.City = "Sukabumi"

	//tidak merubah address1
	address3 = &Address{
		City:     "Sukabumi",
		Province: "West Java",
		Country:  "Cimanggu",
	}

	//merubah semua yang merujuk ke address1
	*address2 = Address{
		City:     "Tasikmalaya",
		Province: "Sulawesi Utara",
		Country:  "Indonesia",
	}

	address5.City = "Jakarta"
	ChangeCountryToEngland(&address6)

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)
	fmt.Println(address5)
	fmt.Println(address6)

}

//pointer function
func ChangeCountryToEngland(address *Address) {
	address.Country = "England"
}
