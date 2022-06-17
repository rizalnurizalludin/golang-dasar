package main

import (
	"fmt"
)

func getFibonacci(n []int) []int {
	if n == 2 {
		return []int{1, 1}
	}
	fib := getFibonacci(n - 1)

	nextFib := fib[len(fib)-1] + fib[len(fib)-2]
	return append(fib, nextFib)
}

func rata_rata(x int, pembagi int) int {
	avv := x / pembagi
	return avv
}

func main() {
	val := getFibonacci(20)
	fmt.Println(val)
	var banyak int
	var rata2 int
	var pembilang int = 0
	fmt.Println("Masukkan Banyak Angka :")
	fmt.Scanln(&banyak)
	var arr = make([]int, banyak)
	for i := 0; i < banyak; i++ {
		fmt.Println("Masukkan Angka :")
		fmt.Scanln(&arr[i])
		pembilang += arr[i]
	}
	rata_rata(pembilang, banyak)
	fmt.Println("Nilai Rata Rata :")
	fmt.Println(arr)
	fmt.Println(rata2)
}

// func getFibonacci(n int) []int {
// 	if n == 2 {
// 		return []int{1, 1}
// 	}
// 	fib := getFibonacci(n - 1)

// 	nextFib := fib[len(fib)-1] + fib[len(fib)-2]
// 	return append(fib, nextFib)
// }
