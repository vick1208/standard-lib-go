package main

import (
	"errors"
	"fmt"
)

var (
	ValidationErr = errors.New("validation error")
	NotFoundErr   = errors.New("not found error")
)

func GetById(id string) (string, error) {
	if id == "" {
		return id, ValidationErr
	}
	if id != "rin" {
		return id, NotFoundErr
	}

	return id, nil
}

func main() {
	id, err := GetById("ri")

	fmt.Println(id)

	if err != nil {
		if errors.Is(err, ValidationErr) {
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundErr) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown error")
		}

	}
}
