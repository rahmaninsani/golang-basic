package main

import (
	"flag"
	"fmt"
)

func main() {
	var username string

	var host *string = flag.String("host", "localhost", "Put your database hostname")
	flag.StringVar(&username, "username", "root", "put your database username")
	password := flag.String("password", "root", "put your database password")

	flag.Parse() // wajib

	fmt.Println("host:", *host)
	fmt.Println("username:", username)
	fmt.Println("password:", *password)

}
