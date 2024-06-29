package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "root", "username database")
	password := flag.String("password", "root", "password database")
	host := flag.String("host", "root", "host database")
	port := flag.Int("port", 0, "host database")

	flag.Parse()

	fmt.Printf("username %s\n", *username)
	fmt.Printf("password %s\n", *password)
	fmt.Printf("host %s\n", *host)
	fmt.Println("port", *port)
}
