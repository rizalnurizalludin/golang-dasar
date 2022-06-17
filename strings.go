package main

import (
	"fmt"
	"strings"
)

func main() {

	//mengecek apakah string mengandung string lain
	fmt.Println(strings.Contains("Rizal Nurizalludin", "Rizal"))
	fmt.Println(strings.Contains("Rizal Nurizalludin", "Hera"))

	//untuk memotong string menjadi slice
	fmt.Println(strings.Split("Rizal Nurizalludin", " "))

	//membuat huruf besar semua dan kecil semua
	fmt.Println(strings.ToLower("Rizal Nurizalludin"))
	fmt.Println(strings.ToUpper("Rizal Nurizalludin"))
	fmt.Println(strings.ToTitle("Rizal Nurizalludin"))

	//menghapus karakter dikiri dan kanan
	fmt.Println(strings.Trim("ssssssssssRizal Nurizalludinssssssssss", "s"))

	//mengganti string dengan string lain
	fmt.Println(strings.ReplaceAll("Rizal Nurizalludin Rizal", "Rizal", "Ganteng"))

}
