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

func (s UserSlice) Len() int           { return len(s) }
func (s UserSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s UserSlice) Less(i, j int) bool { return s[i].Age < s[j].Age }

func main() {
	users := []User{
		{"Victor", 25},
		{"Budi", 26},
		{"Eko", 29},
		{"Adit", 35},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
