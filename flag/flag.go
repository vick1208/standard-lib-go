package main

import (
	"flag"
	"fmt"
)

func main() {
	var username = flag.String("username", "root", "db username")
	var password = flag.String("password", "root", "db password")
	var host = flag.String("host", "localhost", "db host")
	var port = flag.Int("port", 0, "db port")

	flag.Parse()

	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Host", *host)
	fmt.Println("Port", *port)

}
