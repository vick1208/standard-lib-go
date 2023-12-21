package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack(1)
	data.PushBack("Eko")
	data.PushBack("Soegianto")

	head := data.Front()
	fmt.Print(head.Value)
	next := head.Next()
	fmt.Println(next.Value)
	next = next.Next()
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
