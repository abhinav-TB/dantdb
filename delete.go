package dantdb

import (
	"fmt"
	"os"
	"path/filepath"
)

func (d *Driver) DeleteResource(collection, resource string) error {
	mutex := d.getMutex(collection)

	mutex.Lock()
	defer mutex.Unlock()

	path := filepath.Join(collection, resource+".json")

	dir := filepath.Join(d.dir, path)

	if err := os.Remove(dir); err != nil {
		return fmt.Errorf("remove file: %w", err)
	}

	return nil
}

func (d *Driver) DeleteCollection(collection string) error {

	mutex := d.getMutex(collection)

	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, collection)

	if err := os.RemoveAll(dir); err != nil {
		return fmt.Errorf("remove all: %w", err)
	}

	return nil
}
