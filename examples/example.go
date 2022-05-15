package main

import (
	"fmt"

	"github.com/abhinav-TB/dantdb"
)

type Student struct {
	Name   string
	RollNo int
}

func main() {
	dir := "./"

	db, err := dantdb.NewDatabase(dir) // creates a new database
	if err != nil {
		fmt.Println("Error", err)
	}

	db.Write("students", "John", Student{ // writes to the database
		Name:   "John",
		RollNo: 21,
	})

	record := Student{}
	if db.Read("students", "John", &record) != nil { // reads from the database
		fmt.Println("Error", err)
	}
	fmt.Println(record)

}
