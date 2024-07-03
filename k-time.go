package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now.Local())

	var utc time.Time = time.Date(2021, time.April, 12, 0, 0, 0, 0, time.Local)
	fmt.Println(utc)
	fmt.Println(utc.Local().UTC())

	// value to time
	formatter := "2006-01-02 15:04:05"

	value := "2021-12-03 18:28:48"
	valueTime, err := time.Parse(formatter, value)

	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("\n ======== Value to time========\n")

		fmt.Println(valueTime)

		fmt.Println(valueTime.Year())
		fmt.Println(valueTime.Month())
		fmt.Println(valueTime.Day())
		fmt.Println(valueTime.Hour())

	}

}
