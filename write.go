package dantdb

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	extension = ".json"
	temp      = ".tmp"
)

func (d *Driver) Write(collection, resource string, v any) error {
	if collection == "" {
		return ErrNoCollection
	}

	if resource == "" {
		return ErrNoResource
	}

	mutex := d.getMutex(collection)

	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, collection)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("make dir: %w", err)
	}

	b, err := json.MarshalIndent(v, "", "\t")

	if err != nil {
		return fmt.Errorf("marshal data: %w", err)
	}

	dir = filepath.Join(dir, resource+extension)

	// b = append(b, byte('\n'))
	err = os.WriteFile(dir+temp, b, 0600)
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	err = os.Rename(dir+temp, dir)
	if err != nil {
		return fmt.Errorf("rename dir: %w", err)
	}

	return nil
}
