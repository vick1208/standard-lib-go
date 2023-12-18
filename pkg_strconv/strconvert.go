package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("t")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(result)
		fmt.Printf("type = %T \n", result)
	}

	resInt, err := strconv.Atoi("2000")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resInt)
		fmt.Printf("type = %T \n", resInt)
	}

	bin := strconv.FormatInt(999, 2)
	fmt.Println(bin)

	strInt := strconv.Itoa(999)

	fmt.Println(strInt)

}
