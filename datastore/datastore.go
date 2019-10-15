package datastore

import (
	"log"
)

const (
	// INMEMORY storage type
	INMEMORY = iota
	// DISK storage type
	DISK
)

// DataStore generic interface for various data stores
type DataStore interface {
	// Add inserts an entry into the data store
	Add(data interface{}) error
	// Fetch retrieves an entry from data store
	Fetch(size int) []interface{}
}

// StoreInMem place holder for storing lead data
type storeInMem struct {
	inmem []interface{}
}

// Init initializer for the kind of data store
func Init(storageType int) DataStore {
	switch storageType {
	case INMEMORY:
		storage := new(storeInMem)
		return DataStore(storage)
	default:
		log.Println("Invalid-storage-type")
		return nil
	}
}

// Add inserts an entry into the data store
func (ds *storeInMem) Add(data interface{}) error {
	ds.inmem = append(ds.inmem, data)
	return nil
}

// Fetch retrieves an entry from data store
func (ds *storeInMem) Fetch(size int) []interface{} {
	var idx int
	var data []interface{}
	if size == 0 {
		return nil
	}
	if size > len(ds.inmem) {
		data = make([]interface{}, len(ds.inmem))
	} else {
		data = make([]interface{}, size)
	}
	for _, val := range ds.inmem {
		if idx < size {
			data[idx] = val
			idx++
		}
	}
	return data
}
