package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, _ := os.Hostname()

	fmt.Println(hostname)
}
