package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("NOW", now.Local())

	utc := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("UTC", utc.Local())

	parse, err := time.Parse(time.RFC3339, "2016-01-02T15:04:05Z")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(parse)
	}

	fmt.Println(now)
}
