# datastore
--
    import "github.com/anandkolli/Integrate/datastore"


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

#### type StoreInMem

```go
type StoreInMem struct {
}
```

StoreInMem place holder for storing lead data

#### func (*StoreInMem) Add

```go
func (ds *StoreInMem) Add(data interface{}) error
```
Add inserts an entry into the data store

#### func (*StoreInMem) Fetch

```go
func (ds *StoreInMem) Fetch(size int) []interface{}
```
Fetch retrieves an entry from data store
