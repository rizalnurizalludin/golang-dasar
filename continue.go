package main

import "fmt"

func main() {

	/*continue digunakan untuk menghentikan perulangan yang berjalan,
	dan melanjutkan ke perulangan selanjutnya,
	post statement akan tetap di eksekusi*/
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke", i)
	}
}
