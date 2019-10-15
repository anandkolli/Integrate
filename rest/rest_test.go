package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/anandkolli/Integrate/datastore"
)

var (
	client = &http.Client{
		Timeout: time.Second * 10,
	}
	url = "http://demo:demo@localhost:9090" + "/integrate/v1/leaddata"
)

// init function to initialize Integrate server
func init() {
	// create an instance for data storage
	Storage = datastore.Init(datastore.INMEMORY)

	// start integrate rest server
	go StartIntegrateServer()
}

// sendPostReq function to send POST request
func sendPostReq(data LeadData, endpoint string) error {
	var req *http.Request

	// Buffer for request payload
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(data)

	// For testing to pass valid and invalid URLs
	if endpoint == "" {
		req, _ = http.NewRequest("POST", url, buf)
	} else {
		req, _ = http.NewRequest("POST", endpoint, buf)
	}

	// send request to server
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return errors.New("Failure")
	}
	return nil
}

// sendGetReq function to send Get request
func sendGetReq(size string, present bool, endpoint string) []LeadData {
	var req *http.Request

	// For testing to pass valid and invalid URLs
	if endpoint == "" {
		req, _ = http.NewRequest("GET", url, nil)
	} else {
		req, _ = http.NewRequest("GET", endpoint, nil)
	}

	// checks if present is set to true include size parameter in URL else exclude
	if present == true {
		q := req.URL.Query()
		q.Set("size", size)
		req.URL.RawQuery = q.Encode()
	}

	// Send the request
	response, err := client.Do(req)
	if err != nil {
		return nil
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil
	}

	// Read response body
	rsp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil
	}

	// Unmarshal response data to leaddata
	var data []LeadData
	if err := json.Unmarshal(rsp, &data); err != nil {
		return nil
	}

	return data
}

// Test_1 Invalid login credentials
func Test_1(t *testing.T) {

	time.Sleep(1 * time.Second)
	endpoint := "http://dem:demo@localhost:9090" + "/integrate/v1/leaddata"
	payload := LeadData{"anand", "kolli", "akolli@xyz.com", "Integrate", "yes", "TW74DJ", "15-Oct-2019"}
	tests := []struct {
		name string
		args LeadData
		want error
	}{
		{"Invalid-Credentials", payload, errors.New("Failure")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//comparing return values by connect api and expecting return values(wantErr)
			var ret error
			ret = sendPostReq(tt.args, endpoint)
			if ok := reflect.DeepEqual(ret, tt.want); ok == false {
				t.Errorf("error")
			}
		})
	}
}

// Test_2 successful create lead data
func Test_2(t *testing.T) {

	time.Sleep(1 * time.Second)
	payload := LeadData{"anand", "kolli", "akolli@xyz.com", "Integrate", "yes", "TW74DJ", "15-Oct-2019"}
	tests := []struct {
		name string
		args LeadData
		want error
	}{
		{"Successful-Create-LeadData", payload, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//comparing return values by connect api and expecting return values(wantErr)
			if ret := sendPostReq(tt.args, ""); ret != tt.want {
				t.Errorf("error")
			}
		})
	}
}

// Test_3 missing mandatory attribute
func Test_3(t *testing.T) {

	payload := LeadData{"", "kolli", "akolli@xyz.com", "Integrate", "yes", "TW74DJ", "15-Oct-2019"}
	tests := []struct {
		name string
		args LeadData
		want error
	}{
		{"Missing-mandatory-attribute", payload, errors.New("Failure")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//comparing return values by connect api and expecting return values(wantErr)
			var ret error
			ret = sendPostReq(tt.args, "")
			if ok := reflect.DeepEqual(ret, tt.want); ok == false {
				t.Errorf("error")
			}
		})
	}
}

// Test_4 successful get leaddata
func Test_4(t *testing.T) {

	time.Sleep(1 * time.Second)

	data := LeadData{"anand", "kolli", "akolli@xyz.com", "Integrate", "yes", "TW74DJ", "15-Oct-2019"}
	payload := []LeadData{data}
	tests := []struct {
		name string
		args string
		want []LeadData
	}{
		{"Successful-Get-LeadData", "1", payload},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//comparing return values by connect api and expecting return values(wantErr)
			var ret []LeadData
			ret = sendGetReq(tt.args, true, "")
			if ok := reflect.DeepEqual(ret[0], tt.want[0]); ok == false {
				t.Errorf("error")
			}
		})
	}
}

// Test_5 when query parameter size is greater than the stored number of events
func Test_5(t *testing.T) {

	time.Sleep(1 * time.Second)

	data := LeadData{"anand", "kolli", "akolli@xyz.com", "Integrate", "yes", "TW74DJ", "15-Oct-2019"}
	payload := []LeadData{data}
	tests := []struct {
		name string
		args string
		want []LeadData
	}{
		{"Get-When-Size-is-greater-than-storeddata", "2", payload},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//comparing return values by connect api and expecting return values(wantErr)
			var ret []LeadData
			ret = sendGetReq(tt.args, true, "")
			if ok := reflect.DeepEqual(ret[0], tt.want[0]); ok == false {
				t.Errorf("error")
			}
		})
	}
}

// Test_6 when query parameter size is missing from URL
func Test_6(t *testing.T) {

	time.Sleep(1 * time.Second)

	data := LeadData{"anand", "kolli", "akolli@xyz.com", "Integrate", "yes", "TW74DJ", "15-Oct-2019"}
	payload := []LeadData{data}
	tests := []struct {
		name string
		args string
		want []LeadData
	}{
		{"Size-is-missing", "2", payload},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//comparing return values by connect api and expecting return values(wantErr)
			var ret []LeadData
			if ret = sendGetReq(tt.args, false, ""); ret != nil {
				t.Errorf("error")
			}
		})
	}
}

// Test_7 invalid credentials for get lead data
func Test_7(t *testing.T) {

	endpoint := "http://dem:demo@localhost:9090" + "/integrate/v1/leaddata"

	tests := []struct {
		name string
		args string
		want []LeadData
	}{
		{"Invalid-credentials-for-Get-Leaddata", "1", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//comparing return values by connect api and expecting return values(wantErr)
			var ret []LeadData
			if ret = sendGetReq(tt.args, false, endpoint); ret != nil {
				t.Errorf("error")
			}
		})
	}
}

// Test_8 missing user name and password
func Test_8(t *testing.T) {

	endpoint := "http://localhost:9090" + "/integrate/v1/leaddata"

	tests := []struct {
		name string
		args string
		want []LeadData
	}{
		{"Missing-Username-Password", "1", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//comparing return values by connect api and expecting return values(wantErr)
			var ret []LeadData
			if ret = sendGetReq(tt.args, false, endpoint); ret != nil {
				t.Errorf("error")
			}
		})
	}
}
