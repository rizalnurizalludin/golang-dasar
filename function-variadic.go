package main

import "fmt"

func main() {

	total := sumAll(10, 10, 10, 10, 10, 10, 20, 30)
	fmt.Println(total)

	//kode program slice parameters

	slice := []int{100, 200, 300, 400, 500}
	total = sumAll(slice...)
	fmt.Println(total)

}

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}
