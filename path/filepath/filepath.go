package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("hello/package.go"))
	fmt.Println(filepath.Base("hello/package.go"))
	fmt.Println(filepath.Ext("hello/package.go"))
	fmt.Println(filepath.IsAbs("hello/package.go"))
	fmt.Println(filepath.IsLocal("hello/package.go"))
	fmt.Println(filepath.Join("hello", "world", "package.go"))
}
