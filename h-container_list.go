package main

import (
	"fmt"
	"container/list"
)

func main() {
	var data = list.New()

	data.PushBack("O")
	data.PushBack("g")
	data.PushBack("i")

	var head = data.Front()
	fmt.Println(head.Value) //O

	next := head.Next()
	fmt.Println(next.Value) //g

	next = next.Next() 
	fmt.Println(next.Value) //i

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}