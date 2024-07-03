package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"John", "Paul", "Three"}
	values := []int{50, 60, 55}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Eko"))
	fmt.Println(slices.Index(names, "Eko"))
	fmt.Println(slices.Index(names, "Three"))

}
