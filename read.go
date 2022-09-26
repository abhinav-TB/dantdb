package dantdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func (d *Driver) Read(collection, resource string, v any) error {

	if collection == "" {
		return ErrNoCollection
	}

	if resource == "" {
		return ErrNoResource
	}

	record := filepath.Join(d.dir, collection, resource)

	mutex := d.getMutex(collection)
	mutex.Lock()
	b, err := os.ReadFile(record + ".json")
	mutex.Unlock()

	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	return json.Unmarshal(b, &v)
}

func (d *Driver) ReadAll(collection string) ([]string, error) {

	if collection == "" {
		return nil, ErrNoCollection
	}

	mutex := d.getMutex(collection)

	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, collection)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read dir: %w", err)
	}

	records := make([]string, len(files))

	for i, file := range files {
		b, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}

		records[i] = string(b)
	}

	return records, nil
}
