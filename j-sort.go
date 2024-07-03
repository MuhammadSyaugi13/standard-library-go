package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"Eko", 32},
		{"Aji", 22},
		{"Andi", 29},
		{"Udin", 31},
	}

	fmt.Println(UserSlice(users))
	sort.Sort(UserSlice(users))

	fmt.Println(users) // users berubah karena UserSlice mengakses pointer dari struct User

}
