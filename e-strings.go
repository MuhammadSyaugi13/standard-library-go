package main

import (
	"fmt"
	"strings"
)

func main () {

	fmt.Println(strings.Contains("Syaugi", "ugi"))
	fmt.Println(strings.Split("M Syaugi", " "))
	fmt.Println(strings.ToLower("Syaugi"))
	fmt.Println(strings.ToUpper("Syaugi"))
	fmt.Println(strings.Trim("     Syaugi    ", " "))
	fmt.Println(strings.ReplaceAll("M Syaugi", "gi", "qy"))

}