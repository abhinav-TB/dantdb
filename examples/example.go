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

	records, err := db.ReadAll("students") // read all records from database
	if err != nil {
		log.Fatalln(err)
	}

	allusers := []Student{}

	for _, f := range records {
		var student Student
		json.Unmarshal([]byte(f), &student)
		allusers = append(allusers, student)
	}

	fmt.Println(allusers)

	err = db.DeleteResource("students", "John")
	if err != nil { // delete a single document
		log.Fatalln(err)
	}

	err = db.DeleteCollection("students")
	if err != nil { // delete all documents
		log.Fatalln(err)
	}
}
