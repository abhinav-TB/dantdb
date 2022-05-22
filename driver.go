package dantdb

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

type Driver struct {
	mutex   sync.RWMutex
	mutexes map[string]*sync.Mutex
	dir     string
}

func New(dir string) (*Driver, error) {
	dir = filepath.Clean(dir)

	_, err := os.Stat(dir)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				return nil, fmt.Errorf("make dir: %w", err)
			}
		} else {
			return nil, fmt.Errorf("get stats: %w", err)
		}
	}

	return &Driver{
		dir:     dir,
		mutexes: make(map[string]*sync.Mutex),
	}, nil
}
