package main

import "fmt"

func main() {

	runApplication()

}

func logging() {
	fmt.Println("Selesai Memanggil Function")
}

func runApplication() {

	/*//-defer adalah function yang bisa dijadwalkan untuk
	dieksekusi setelah sebuah function selesai dieksekusi
	//-defer akan selalu dieksekusi walaupun terjadi error
	di function yang dieksekusi*/
	defer logging()
	fmt.Println("Run Application")

}
