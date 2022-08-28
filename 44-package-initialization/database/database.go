package database

import "fmt"

var databaseName string

func init() {
	fmt.Println("Initializing database")
	databaseName = "MySQL"
}

func GetDatabaseName() string {
	return databaseName
}
