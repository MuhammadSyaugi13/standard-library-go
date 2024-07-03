package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5)

  /*dengan perulangan*/ 
	for i:= 1; i< (data.Len()+1); i++ {
		data.Value = "Value "+ strconv.Itoa(i)
		data = data.Next()
	}
  /*./ dengan perulangan*/ 

	// data.Value = "value 1"
	
	// data = data.Next()
	// data.Value = "value 2"
	
	// data = data.Next()
	// data.Value = "value 3"
	
	// data = data.Next()
	// data.Value = "value 4"

	// data = data.Next()
	// data.Value = "value 5"
	
	// // data = data.Next()
	// // data.Value = "value 6"
	

	data.Do(func (value any) {
		fmt.Println(value)
	})
}