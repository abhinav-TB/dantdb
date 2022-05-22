package dantdb

import "sync"

func (d *Driver) getMutex(collection string) *sync.Mutex {

	d.mutex.RLock()
	m, ok := d.mutexes[collection]
	d.mutex.RUnlock()

	if !ok {
		m = &sync.Mutex{}

		d.mutex.Lock()
		d.mutexes[collection] = m
		d.mutex.Unlock()
	}

	return m
}
