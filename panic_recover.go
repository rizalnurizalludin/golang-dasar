package main

import "fmt"

func main() {

	runApp(true)
	fmt.Println("Rizal Nurizalludin")

}

func endApp() {
	/**
	-recover adalah function yang
	digunakan untuk menangkap data panic
	-dengan recover proses panic
	akan berhenti sehingga program akan tetap berjalan
	-recover disimpan didalam defer function
	*/
	message := recover()
	if message != nil {
		fmt.Println("Aplikasi Error dengan Message", message)
	}

	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		/*panic adalah function yang bisa digunakan
		untuk menghentikan program
		-biasa dipangging ketika error pada saat program berjalan
		-namun defer akan tetap dieksekusi*/
		panic("Error")
	}
	fmt.Println("Start App")
}
