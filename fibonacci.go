package main

import (
	"fmt"
)

func getFibonacci(n int) []int {
	if n == 2 {
		return []int{1, 1}
	}
	fib := getFibonacci(n - 1)

	nextFib := fib[len(fib)-1] + fib[len(fib)-2]
	return append(fib, nextFib)
}
func main() {
	val := getFibonacci(20)
	fmt.Println(val)
}
