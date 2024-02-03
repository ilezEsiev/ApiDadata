package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddressSearchHandler(t *testing.T) {
	requestBody := `{"query": "SomeAddress"}`
	req, err := http.NewRequest("POST", "/api/search/address", bytes.NewBufferString(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	AddressSearchHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedResponseBody := `{"Addresses":[]}`
	if rr.Body.String() != expectedResponseBody {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expectedResponseBody)
	}
}

func TestAddressGeocodeHandler(t *testing.T) {
	requestBody := `{"address": "SomeAddress"}`
	req, err := http.NewRequest("POST", "/api/geocode/address", bytes.NewBufferString(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	AddressGeocodeHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseMap map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
	}

	if _, ok := responseMap["someExpectedField"]; !ok {
		t.Errorf("Handler returned unexpected response: %v", responseMap)
	}
}
