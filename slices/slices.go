package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Fred", "Velma", "Shaggy", "Daphne"}
	numVal := []int{100, 90, 98, 95, 96, 80}

	fmt.Println(slices.Min(numVal))
	fmt.Println(slices.Contains(names, "Fred"))
	fmt.Println(slices.Index(names, "Fred"))
}
