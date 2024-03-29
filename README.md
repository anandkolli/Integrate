# Integrate Assesment

## Pre requisite
- Install "github.com/gorilla/mux"
- Tested with go version go1.13.1 linux/amd64
- Ubunut 16.04

## Implementation details
1) Integrate server is running on localhost:9090
2) One REST endpoint is exposed (/integrate/v1/leaddata) and two operations can be performed on this endpoint
    
    a) POST => To send the lead data to server
    
    b) GET => To get the lead data from the server (I have added size parameter to request the number of lead data entries from server)

3) Basic authentication is required to proceed with the above endpoint.

    a) Provide username/password along with the URL to proceed with the operation
    
    b) Username and password is configured to demo/demo

4) Data is stored inmemory and implementation can be enhanced 
```
curl -v -H "Content-Type: application/json" -XPOST http://demo:demo@localhost:9090/integrate/v1/leaddata -d @data.json

data.json
{
  "data": {
    "firstname": "anand",
    "lastname": "kolli",
    "email": "akolli@xyz.com",
    "company": "integrate",
    "acceptterms": "yess",
    "postcode": "TW74DJ",
    "date": "16-Oct-2019"
  }
}

```
```
curl -v -H "Content-Type: application/json" -XGET http://demo:demo@localhost:9090/integrate/v1/leaddata?size=1
```

5) Used gofmt, goimports and golint for formatting and static code analysis.
 
## How to Run
- Clone the repository and run below commands
```
cd rest
go test -v -cover
=== RUN   Test_1
2019/10/16 13:26:01 Spinning-Integrate-Server
=== RUN   Test_1/Invalid-Credentials
--- PASS: Test_1 (1.08s)
    --- PASS: Test_1/Invalid-Credentials (0.00s)
=== RUN   Test_2
=== RUN   Test_2/Successful-Create-LeadData
--- PASS: Test_2 (1.08s)
    --- PASS: Test_2/Successful-Create-LeadData (0.00s)
=== RUN   Test_3
=== RUN   Test_3/Missing-mandatory-attribute
2019/10/16 13:26:03 LeadData-validation-failed Missing-Mandatory-Attribute
--- PASS: Test_3 (0.00s)
    --- PASS: Test_3/Missing-mandatory-attribute (0.00s)
=== RUN   Test_4
=== RUN   Test_4/Successful-Get-LeadData
--- PASS: Test_4 (1.01s)
    --- PASS: Test_4/Successful-Get-LeadData (0.00s)
=== RUN   Test_5
=== RUN   Test_5/Get-When-Size-is-greater-than-storeddata
--- PASS: Test_5 (1.04s)
    --- PASS: Test_5/Get-When-Size-is-greater-than-storeddata (0.00s)
=== RUN   Test_6
=== RUN   Test_6/Size-is-missing
2019/10/16 13:26:06 Bad-request-size-missing
--- PASS: Test_6 (1.01s)
    --- PASS: Test_6/Size-is-missing (0.00s)
=== RUN   Test_7
=== RUN   Test_7/Invalid-credentials-for-Get-Leaddata
--- PASS: Test_7 (0.00s)
    --- PASS: Test_7/Invalid-credentials-for-Get-Leaddata (0.00s)
=== RUN   Test_8
=== RUN   Test_8/Missing-Username-Password
--- PASS: Test_8 (0.00s)
    --- PASS: Test_8/Missing-Username-Password (0.00s)
PASS
coverage: 93.3% of statements
ok      github.com/anandkolli/Integrate/rest    5.233s
```

## Future Scope
1) Can add couple of more api's for login and logout along with JWT tokens.
   - Login API is commented out for the sake of simplicity.

2) Design of datastore package is extensible and can push event data to persistent storage as well.
