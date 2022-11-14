package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/abhinav-TB/dantdb"
)

type Student struct {
	Name   string
	RollNo int
}

func main() {
	dir := "./"

	db, err := dantdb.New(dir) // creates new database
	if err != nil {
		log.Fatalln(err)
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

	filter := []string{"Paul", "Robert", "Albert"} // read filtered records from database
	filtered, err := db.ReadFiltered("students", filter)
	if err != nil {
		log.Fatalln(err)
	}

	filteredUsers := []Student{}

	for _, f := range filtered {
		var student Student
		json.Unmarshal([]byte(f), &student)
		filteredUsers = append(filteredUsers, student)
	}

	fmt.Printf("Filtered %d of expected %d students: %v\n", len(filteredUsers), len(filter), filteredUsers)

	records, err := db.ReadAll("students") // read all records from database
	if err != nil {
		log.Fatalln(err)
	}

	allUsers := []Student{}

	for _, f := range records {
		var student Student
		json.Unmarshal([]byte(f), &student)
		allUsers = append(allUsers, student)
	}

	fmt.Printf("All users: %v\n", allUsers)

	err = db.DeleteResource("students", "John")
	if err != nil { // delete a single document
		log.Fatalln(err)
	}

	err = db.DeleteCollection("students")
	if err != nil { // delete all documents
		log.Fatalln(err)
	}
}
