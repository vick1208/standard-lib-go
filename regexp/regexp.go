package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`E([a-z])o`)

	fmt.Println(regex.MatchString("Eko"))
	fmt.Println(regex.MatchString("Elo"))
	fmt.Println(regex.MatchString("EAo"))
	fmt.Println(regex.MatchString("Edo"))

	fmt.Println(regex.FindAllString("Eko E2o Eto Elo EKo EKO", 10))

}
