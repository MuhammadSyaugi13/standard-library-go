package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "eko kurniawan"

	fmt.Println([]byte(value))

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("error: ", err.Error())
	} else {
		fmt.Println(string(decoded))
	}

}
