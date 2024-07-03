package main

import (
	"fmt"
	"time"
)

func main() {
	var duration1 = 90 * time.Second
	var duration2 = 6 * time.Minute
	var duration3 = duration2 - duration1

	fmt.Println(duration3)
	fmt.Printf("duration : %d\n", duration3)

}
