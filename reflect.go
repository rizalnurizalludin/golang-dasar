package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
	Age  int
}

type Book struct {
	Name        string `required:"true"` //tag
	Description string `required:"true"`
}

func main() {
	sample := Sample{"Rizal", 26}
	sampleType := reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(1).Name)

	//tag
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	sample.Name = ""
	fmt.Println(IsValid(sample))

	book := Book{"Naruto", "Anime Comic"}
	fmt.Println(IsValid(book))

}

func IsValid(data interface{}) bool {

	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}
