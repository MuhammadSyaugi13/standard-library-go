package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	ValidationError = errors.New("Error when validation")
	NotFoundError   = errors.New("Error Not Found")
)

func getById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "ogi" {
		return NotFoundError
	}

	return nil
}

func getArgs() string {
	argument := os.Args

	fmt.Println(argument)

	if len(argument) < 2 {
		return ""
	}

	return argument[1]
}

func main() {
	argument := getArgs()

	err := getById(argument)

	if errors.Is(err, ValidationError) {
		fmt.Println(ValidationError)
	}

	if errors.Is(err, NotFoundError) {
		fmt.Println(NotFoundError)
	}
	if err == nil {
		fmt.Println("data Ditemukan")
	}
}
