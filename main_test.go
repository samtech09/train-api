package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetByIDApi(t *testing.T) {
	req, err := http.NewRequest("GET", "/getrecordbyid", nil)
	if err != nil {
		t.Fatal()
	}

	// add query parameter
	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()

	// create response recorder to read statuscode and body
	rr := httptest.NewRecorder()

	// create handler and serve endpoint
	handler := http.HandlerFunc(getbyid)
	handler.ServeHTTP(rr, req)

	// check http status
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect
	expected := `{"ID":1,"Title":"Mango","Price":45.55}` + "\n"
	rsp := rr.Body.String()
	if rsp != expected {
		t.Errorf("handler returned unexpected response: \n\tgot  %v \n\twant %v", rsp, expected)
	}
}
