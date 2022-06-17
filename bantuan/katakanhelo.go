package bantuan

//nama function harus besar jika ingin diakses dari luar package

//bisa diakses package lain
var Application = "golang"

func Katakanhelo(n string) string {
	return "Hello " + n
}

//tidak bisa diakses package lain
var version = "1.0.0"

func katakanhelo(s string) string {
	return "Helo " + s
}
