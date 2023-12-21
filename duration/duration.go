package main

import (
	"fmt"
	"time"
)

func main() {
	dur1 := 100 * time.Second
	dur2 := 10 * time.Millisecond
	dur3 := dur1 - dur2

	fmt.Println(dur3)
	fmt.Printf("duration: %d\n", dur3)
}
