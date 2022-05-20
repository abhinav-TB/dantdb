package main

import (
	"encoding/json"
	"fmt"

	"github.com/abhinav-TB/dantdb"
)

type Student struct {
	Name   string
	RollNo int
}

func main() {
	dir := "./"

	db, err := dantdb.NewDatabase(dir) // creates new database
	if err != nil {
		fmt.Println("Error", err)
	}

	students := []Student{
		{"John", 1},
		{"Paul", 2},
		{"Robert", 3},
		{"Vince", 4},
		{"Neo", 5},
		{"Albert", 6},
	}

	for _, value := range students { // write to database
		db.Write("students", value.Name, Student{
			Name:   value.Name,
			RollNo: value.RollNo,
		})
	}

	records, err := db.ReadAll("students") // read all records from database
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

	if err := db.Delete("students", "John"); err != nil { // delete a single document
		fmt.Println("Error", err)
	}

	if err := db.Delete("students", ""); err != nil { // delete all documents
		fmt.Println("Error", err)
	}
}
