package main

import (
	"encoding/json"
	"fmt"

	"github.com/abhinav-TB/datantdb"
)

type Student struct {
	Name   string
	RollNo int
}

func main() {
	dir := "./"

	db, err := datantdb.NewDatabase(dir) // creates a new database
	if err != nil {
		fmt.Println("Error", err)
	}

	students := []Student{
		{"John", 1},
		{"Paul", 2},
		{"Robert", 3},
	}

	// writes the data to the database
	for _, value := range students {
		db.Write("students", value.Name, Student{
			Name:   value.Name,
			RollNo: value.RollNo,
		})
	}

	// reads all the records from the database
	records, err := db.ReadAll("students")
	if err != nil {
		fmt.Println("Error", err)
	}

	allusers := []Student{}

	for _, f := range records {
		var student Student
		json.Unmarshal([]byte(f), &student)
		allusers = append(allusers, student)
	}

	fmt.Println(allusers)

	// reads a single record from the database
	record := Student{}
	if db.Read("students", "Paul", &record) != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(record)

	// deletes single record from the database
	if err := db.Delete("students", "John"); err != nil {
		fmt.Println("Error", err)
	}

	// deletes all the records from the database
	if err := db.Delete("students", ""); err != nil {
		fmt.Println("Error", err)
	}
}
