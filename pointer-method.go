package main

import "fmt"

type Man struct {
	Name string
}

func main() {

	guy1 := Man{
		Name: "Rizal",
	}
	guy1.Married()
	fmt.Println(guy1.Name)

}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
	fmt.Println("Di Method", man.Name)
}
