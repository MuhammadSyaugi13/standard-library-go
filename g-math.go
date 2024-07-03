package main

import (
	"math"
	"fmt"
)

func main() {

	fmt.Println(math.Ceil(10.2))
	fmt.Println(math.Floor(10.2))
	fmt.Println(math.Round(10.6))
	fmt.Println(math.Round(10.4))
	fmt.Println(math.Max(10, 2))
	fmt.Println(math.Min(10, 2))

	// log 2
	hasilLog := math.Log2(16)
	fmt.Println(hasilLog)
}