package database

import "fmt"

var connection string

func init() {
	fmt.Println("Init Dipanggil")
	connection = "mySQL"
}

func GetDatabase() string {
	return connection
}
