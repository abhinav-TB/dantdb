package dantdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	ErrNoCollection = errors.New("missing collection")
	ErrNoResource   = errors.New("missing resource")
)

func (d *Driver) Read(collection, resource string, v any) error {

	if collection == "" {
		return ErrNoCollection
	}

	if resource == "" {
		return ErrNoResource
	}

	record := filepath.Join(d.dir, collection, resource)

	b, err := os.ReadFile(record + ".json")
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	return json.Unmarshal(b, &v)
}

func (d *Driver) ReadAll(collection string) ([]string, error) {

	if collection == "" {
		return nil, ErrNoCollection
	}

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
