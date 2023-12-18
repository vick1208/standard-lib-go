package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	for _, arg := range args {
		fmt.Println(arg)
	}

	host, err := os.Hostname()

	if err == nil {
		fmt.Println(host)
	} else {
		fmt.Println("Error:", err.Error())
	}

}
