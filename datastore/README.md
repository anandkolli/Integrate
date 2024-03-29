# datastore
--
    import "github.com/anandkolli/Integrate/datastore"

Package datastore is an abstraction for the storage of event data either to
persistent storage or in memory

## Usage

```go
const (
	// INMEMORY storage type
	INMEMORY = iota
	// DISK storage type
	DISK
)
```

#### type DataStore

```go
type DataStore interface {
	// Add inserts an entry into the data store
	Add(data interface{}) error
	// Fetch retrieves an entry from data store
	Fetch(size int) []interface{}
}
```

DataStore generic interface for various data stores

#### func  Init

```go
func Init(storageType int) DataStore
```
Init initializer for the kind of data store
