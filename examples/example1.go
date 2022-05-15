package main

import (
	"encoding/json"
	"fmt"

	"github.com/abhinav-TB/datantdb"
)

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	dir := "./"

	db, err := datantdb.NewDatabase(dir)
	if err != nil {
		fmt.Println("Error", err)
	}

	employees := []User{
		{"John", "23", "23344333", "Myrl Tech", Address{"bangalore", "karnataka", "india", "410013"}},
		{"Paul", "25", "23344333", "Google", Address{"san francisco", "california", "USA", "410013"}},
		{"Robert", "27", "23344333", "Microsoft", Address{"bangalore", "karnataka", "india", "410013"}},
		{"Vince", "29", "23344333", "Facebook", Address{"bangalore", "karnataka", "india", "410013"}},
		{"Neo", "31", "23344333", "Remote-Teams", Address{"bangalore", "karnataka", "india", "410013"}},
		{"Albert", "32", "23344333", "Dominate", Address{"bangalore", "karnataka", "india", "410013"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	// records, err := db.ReadAll("users")
	// if err != nil {
	// 	fmt.Println("Error", err)
	// }
	// fmt.Println(records)

	// allusers := []User{}

	// for _, f := range records {
	// 	employeeFound := User{}
	// 	if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
	// 		fmt.Println("Error", err)
	// 	}
	// 	allusers = append(allusers, employeeFound)
	// }
	// fmt.Println((allusers))

	// Read single document from the database
	record := User{}
	if db.Read("users", "Neo", &record) != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(record)

	// if err := db.Delete("users", "John"); err != nil {
	// 	fmt.Println("Error", err)
	// }

	// if err := db.Delete("users", ""); err != nil {
	// 	fmt.Println("Error", err)
	// }
}
