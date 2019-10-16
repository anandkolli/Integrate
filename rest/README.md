# rest
--
    import "github.com/anandkolli/Integrate/rest"

Package rest package exposing integration rest endpoint to send and collect from
the server

## Usage

```go
const (
	// IntPort default port for Integrate server
	IntPort = "9090"
	// UserName default user to access Integrate services
	UserName = "demo"
	// Password default password to access Integrate services
	Password = "demo"
)
```

```go
var (
	// Storage placeholder for the type of storage used to store Leaddata
	Storage datastore.DataStore
)
```

#### func  StartIntegrateServer

```go
func StartIntegrateServer()
```
StartIntegrateServer starts the REST Server using the handler that is provided

#### type Attributes

```go
type Attributes struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	Company     string `json:"company"`
	AcceptTerms string `json:"acceptterms"`
	PostCode    string `json:"postcode"`
	Date        string `json:"date"`
}
```

Attributes holds attributes of lead data

#### type LeadData

```go
type LeadData struct {
	Data Attributes `json:"data"`
}
```

LeadData holds data collected from events

#### type LoginPayload

```go
type LoginPayload struct {
	// Username in request
	Username string `json:"username"`
	// Password in request
	Password string `json:"password"`
}
```

LoginPayload datatype for handling username/password in the login request
