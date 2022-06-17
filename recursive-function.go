package main

import "fmt"

func main() {

	fmt.Println(factorialLoop(5))

	fmt.Println(factorialLoopRecursive(7))

	for i := 1; i < 10; i++ {
		fmt.Println(fibonacciLoop(i))
	}

}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialLoopRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialLoopRecursive(value-1)
	}
}

func fibonacciLoop(value int) int {
	if value == 1 || value == 2 {
		return value
	} else {
		return fibonacciLoop(value-2) + fibonacciLoop(value-1)
	}
}
