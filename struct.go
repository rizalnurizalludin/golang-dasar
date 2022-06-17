package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	Marriage      bool
}

func main() {

	var cust1 Customer
	cust1.Name = "Rizal"
	cust1.Address = "Indonesia"
	cust1.Age = 25

	fmt.Println(cust1)
	fmt.Println(cust1.Age)

	//Struct Literals

	cust2 := Customer{
		Name:    "Hera",
		Address: "Tasikmalaya",
		Age:     20,
	}

	fmt.Println(cust2)
	fmt.Println(cust2.Age)

	cust3 := Customer{"Budi", "Jakarta", 35, false}

	fmt.Println(cust3)
	fmt.Println(cust3.Age)

	sayHi(cust1, "Ebon")

	//struct method
	cust1.sayHii("Ebons")

}

func sayHi(cust Customer, name string) {
	fmt.Println("Hello", name, "my name is", cust.Name)
}

//struct method
func (cust Customer) sayHii(name string) {
	fmt.Println("Hello", name, "my name is", cust.Name)

}
