package main

import (
	"44-package-initialization/database"
	//_ "44-package-initialization/database"
	"fmt"
)

func main() {
	databaseName := database.GetDatabaseName()
	fmt.Println(databaseName)
}
