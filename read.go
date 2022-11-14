package dantdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

func (d *Driver) ReadFiltered(collection string, resources []string) ([]string, error) {

	if collection == "" {
		return nil, ErrNoCollection
	}

	// Build a set, values are ignored.
	resourceSet := make(map[string]bool)
	for _, r := range resources {
		resourceSet[r] = true
	}

	mutex := d.getMutex(collection)

	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, collection)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read dir: %w", err)
	}

	var records []string

	for _, file := range files {

		// Read only files filtered through a resourceSet.
		r := strings.TrimSuffix(file.Name(), ".json")
		if _, ok := resourceSet[r]; ok {
			b, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, err
			}
			records = append(records, string(b))
		}
	}

	return records, nil
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
