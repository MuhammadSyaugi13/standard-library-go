package main

import (
	"fmt"
	"reflect"
)

type sample struct {
	Name string
}

type person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readfile(value any) {
	var valueType = reflect.TypeOf(value)
	fmt.Println(valueType.Name())
	fmt.Println(valueType.NumField())
	fmt.Println("-----------------")

	for i := 0; i < valueType.NumField(); i++ {
		var structField = valueType.Field(i)
		fmt.Println(structField.Name, "with type ", structField.Type)

		fmt.Println(structField.Tag.Get("required"))
	}

	fmt.Println("=====================\n")
}

func isValid(data interface{}) (result bool) {

	fmt.Println("\n=============================================\n")

	result = true
	t := reflect.TypeOf(data)

	fmt.Print("data : ")
	fmt.Println(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			fmt.Println(reflect.ValueOf(data).Field(i).Interface())

			// data := reflect.ValueOf(data).Field(i).Interface()

			// result = data != ""
			if result == false {
				return result
			}
		}
	}

	return result

}

func main() {

	ogi := sample{"ogi"}
	readfile(ogi)
	readfile(person{"andi", "Jakarta", "andi@gmail.com"})

	valid := isValid(person{"andi", "d", "asd"})
	fmt.Println(valid)

}
