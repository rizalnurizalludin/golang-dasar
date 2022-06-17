package main

import "fmt"

type HasName interface {
	GetName() string
}

type person struct {
	Name string
}

type animal struct {
	Name string
}

func main() {

	var cust1 person
	cust1.Name = "Rizal"
	sayHelo(cust1)

	cat := animal{
		Name: "Push",
	}
	sayHelo(cat)
}

func (Person person) GetName() string {
	return Person.Name
}

func (Animal animal) GetName() string {
	return Animal.Name
}

func sayHelo(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}
