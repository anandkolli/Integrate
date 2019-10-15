# Integrate
Integrate assesment

## Implementation details
1) Integrate server is running on localhost:9090
2) One REST endpoint is exposed (/integrate/v1/leaddata) and two operations can be performed on this endpoint
    
    a) To send the lead data to server
    
    b) To get the lead data from the server (I have added size parameter to request the number of lead data entries from server)

3) Basic authentication is required to proceed with the above endpoint.

    a) Provide username/password along with the URL to proceed with the operation
    
    b) Username and password is configured to demo/demo

4) Data is stored inmemory and implementation can be enhanced 
  ```
curl -v -H "Content-Type: application/json" -XPOST http://demo:demo@localhost:9090/integrate/v1/leaddata -d '{"firstname":"anand","lastname":"kolli","company":"integrate","postcode":"TW74DJ","email":"akolli@xyz.com","acceptterms":"yess","dat":"15-Oct-2019"}'

curl -v -H "Content-Type: application/json" -XGET http://demo:demo@localhost:9090/integrate/v1/leaddata?size=1
  ```

5) Used gofmt, goimports and golint for formatting and static code analysis.
 
## How to Run
- Clone the repository and run below commands
```
cd rest
go test -v

=== RUN   Test_1
2019/10/15 23:19:30 Spinning-Integrate-Server
=== RUN   Test_1/Invalid-Credentials
--- PASS: Test_1 (1.00s)
    --- PASS: Test_1/Invalid-Credentials (0.00s)
=== RUN   Test_2
=== RUN   Test_2/Successful-Create-LeadData
--- PASS: Test_2 (1.01s)
    --- PASS: Test_2/Successful-Create-LeadData (0.00s)
=== RUN   Test_3
=== RUN   Test_3/Missing-mandatory-attribute
2019/10/15 23:19:32 LeadData-validation-failed Missing-Mandatory-Attribute
--- PASS: Test_3 (0.00s)
    --- PASS: Test_3/Missing-mandatory-attribute (0.00s)
=== RUN   Test_4
=== RUN   Test_4/Successful-Get-LeadData
2019/10/15 23:19:33 [{anand kolli akolli@xyz.com Integrate yes TW74DJ 15-Oct-2019}]
--- PASS: Test_4 (1.00s)
    --- PASS: Test_4/Successful-Get-LeadData (0.00s)
=== RUN   Test_5
=== RUN   Test_5/Get-When-Size-is-greater-than-storeddata
2019/10/15 23:19:34 [{anand kolli akolli@xyz.com Integrate yes TW74DJ 15-Oct-2019}]
--- PASS: Test_5 (1.00s)
    --- PASS: Test_5/Get-When-Size-is-greater-than-storeddata (0.00s)
=== RUN   Test_6
=== RUN   Test_6/Size-is-missing
2019/10/15 23:19:35 Bad-request-size-missing
--- PASS: Test_6 (1.00s)
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
ok      github.com/anandkolli/Integrate/rest    5.024s

```

## Future Scope
1) Can add couple of more api's for login and logout along with JWT tokens.
   - Login API is commented out for the sake of simplicity.

2) Can enhance data store package to push the data to persistent storage.
