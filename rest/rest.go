package rest

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/anandkolli/Integrate/datastore"
	"github.com/gorilla/mux"
)

const (
	// IntPort default port for Integrate server
	IntPort = "9090"
	// UserName default user to access Integrate services
	UserName = "demo"
	// Password default password to access Integrate services
	Password = "demo"
)

var (
	isUserLoggedIn = false
	// Storage placeholder for the type of storage used to store Leaddata
	Storage datastore.DataStore
)

// LeadData place holder for data collected from events
type LeadData struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	Company     string `json:"company"`
	AcceptTerms string `json:"acceptterms"`
	PostCode    string `json:"postcode"`
	Date        string `json:"date"`
}

//LoginPayload datatype for handling username/password in the login request
type LoginPayload struct {
	// Username in request
	Username string `json:"username"`
	// Password in request
	Password string `json:"password"`
}

// StartIntegrateServer starts the REST Server using the handler that is provided
func StartIntegrateServer() {
	log.Println("Spinning-Integrate-Server")

	router := mux.NewRouter()
	// Endpoint to login to Integrate interface
	//router.HandleFunc("/integrate/v1/login", Login).Methods("POST")
	router.HandleFunc("/integrate/v1/leaddata", createLeadData).Methods("POST")
	router.HandleFunc("/integrate/v1/leaddata", getLeadData).Methods("GET")

	err := http.ListenAndServe(fmt.Sprintf(":%s", IntPort), router)
	if err != nil {
		log.Println("Error-in-REST-Server", err.Error())
	}
}

// validateLogin Validates payload in login request
func validateLogin(payload LoginPayload) bool {

	if payload.Username == UserName && payload.Password == Password {
		return true
	}
	return false
}

// getCredentials retrieves username and password from URL
func getCredentials(req *http.Request) []string {
	pair := make([]string, 2)
	auth := strings.SplitN(req.Header.Get("Authorization"), " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		return pair
	}

	payloa, _ := base64.StdEncoding.DecodeString(auth[1])
	pair = strings.SplitN(string(payloa), ":", 2)
	return pair
}

// validateLeadData validates leaddata attributes
func validateLeadData(data *LeadData) error {
	if data.FirstName == "" || data.LastName == "" || data.Email == "" || data.AcceptTerms == "" {
		return errors.New("Missing-Mandatory-Attribute")
	}
	return nil
}

// createLeadData endpoint creates an entry into leaddata storage
func createLeadData(rsp http.ResponseWriter, req *http.Request) {
	var payload LeadData

	// validate user credentials
	pair := getCredentials(req)
	if ret := validateLogin(LoginPayload{pair[0], pair[1]}); ret == false {
		rsp.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
		rsp.WriteHeader(http.StatusBadRequest)
		return
	}

	if ret := validateLeadData(&payload); ret != nil {
		log.Println("LeadData-validation-failed",ret)
		rsp.WriteHeader(http.StatusBadRequest)
		return
	}

	Storage.Add(payload)
}

// getLeadData handler to get lead data
func getLeadData(rsp http.ResponseWriter, req *http.Request) {
	pair := getCredentials(req)
	if ret := validateLogin(LoginPayload{pair[0], pair[1]}); ret == false {
		rsp.WriteHeader(http.StatusUnauthorized)
		return
	}

	size, ok := req.URL.Query()["size"]
	if !ok {
		log.Println("Bad-request-size-missing")
		rsp.WriteHeader(http.StatusBadRequest)
		return
	}
	sz, _ := strconv.Atoi(size[0])
	data := Storage.Fetch(sz)
	log.Println(data)
	json.NewEncoder(rsp).Encode(data)
}

/*func Login(rsp http.ResponseWriter, req *http.Request) {
	log.Println("Login-request-received")
	var payload LoginPayload
	//Get details for the payload
	if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
		log.Println("Login-failed-due-bad-request")
		rsp.WriteHeader(http.StatusBadRequest)
	}
	if !validateLogin(payload) {
		log.Println("Login-failed-due-to-invalid-credentials")
		rsp.WriteHeader(http.StatusUnauthorized)
		return
	}
	isUserLoggedIn = true
	rsp.WriteHeader(http.StatusOK)
	log.Println("Login-successful")
}*/
