package main

import (
	"fmt"
	"strconv"
)

func main(){

	result,err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error:", err.Error())
	}else {
		fmt.Println(result)
	}

	resultInt, err := strconv.Atoi("2000")
	if err != nil {
		fmt.Println("Error:", err.Error())
	}else {
		fmt.Println(resultInt)
		fmt.Println(resultInt+1000)
	}

	binary := strconv.FormatInt(5, 2)
	fmt.Println(binary)

	var stringInt string = strconv.Itoa(999)
	fmt.Println(stringInt)

}