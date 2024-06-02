package main

import (
	"errors"
	"fmt"
)

var (
	ValidationErr = errors.New("validation error")
	NotFoundErr   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationErr
	}
	if id != "eko" {
		return NotFoundErr
	}

	return nil
}

func main() {
	err := GetById("budi")

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
