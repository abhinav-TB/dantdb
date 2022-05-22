# dantDB

A simple no-SQL database inspired by SQLite .

## Features

- Easy to Install as a library
- Easy to use
- Easy to visualize using JSON
- safe DB access through mutex's

## Installation

 ```sh
 go get github.com/abhinav-TB/dantdb
 ```

## Usage

```go
package main

import (
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

	db, err := dantdb.New(dir) // creates a new database
	if err != nil {
		log.Fatal(err)
	}

	db.Write("students", "John", Student{ // writes to the database
		Name:   "John",
		RollNo: 21,
	})

	record := Student{}

	err = db.Read("students", "John", &record)

	if err != nil { // reads from the database
		log.Fatal(err)
	}

	fmt.Println(record)

}
```

More examples can be found in the [examples](https://github.com/abhinav-TB/datantdb/tree/master/examples) directory

## Contribute

Contributions are what makes the open source community such an amazing place to learn, inspire, and create. Any
contributions you make are greatly appreciated.

1. Fork the Project
2. Create your Feature Branch (git checkout -b feature/AmazingFeature)
3. Commit your Changes (git commit -m 'Add some AmazingFeature')
4. Push to the Branch (git push origin feature/AmazingFeature)
5. Open a Pull Request

### Top contributors

| [IlliaFox](https://github.com/illiafox) |
|:----------------------------------------|

## License

MIT © [Abhinav TB ](https://github.com/abhinav-TB)
