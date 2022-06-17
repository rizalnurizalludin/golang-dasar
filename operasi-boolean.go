package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusnilaiAkhir = nilaiAkhir > 80
	var lulusAbsensi = absensi > 80
	fmt.Println(lulusnilaiAkhir)
	fmt.Println(lulusAbsensi)

	var lulus = lulusnilaiAkhir && lulusAbsensi
	fmt.Println(lulus)

	fmt.Println(nilaiAkhir > 80 && absensi >= 80)
}
