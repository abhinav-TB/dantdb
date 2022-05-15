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

	db, err := datantdb.NewDatabase(dir)
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

	for _, value := range students {
		db.Write("users", value.Name, Student{
			Name:   value.Name,
			RollNo: value.RollNo,
		})
	}

	records, err := db.ReadAll("users")
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
}
