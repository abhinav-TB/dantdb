package dantdb

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func (d *Driver) Write(collection, resource string, v any) error {
	if collection == "" {
		return ErrNoCollection
	}

	if resource == "" {
		return ErrNoResource
	}

	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return fmt.Errorf("marshal data: %w", err)
	}

	mutex := d.getMutex(collection)

	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, collection)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("make dir: %w", err)
	}

	f := filepath.Join(dir, resource+extension)

	err = os.WriteFile(f, b, 0600)
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}
