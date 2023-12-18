package main

import "fmt"

func main() {
	firstName := "Rian"
	lastName := "Handhoko"

	fmt.Println("Hello '", firstName, lastName, "'")
	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}
