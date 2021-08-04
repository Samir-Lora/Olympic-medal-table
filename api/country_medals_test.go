package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCountry_medals(t *testing.T) {
	// Create a request
	req, err := http.NewRequest("GET", "/api/countries", nil)

	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	rec := httptest.NewRecorder()

	// Check the status code
	if rec.Result().StatusCode != http.StatusOK {
		t.Fatalf("Error : %v$", rec.Result().StatusCode)
	}

	//We Read the response body on the line below.
	j := req.Header.Get("Content-Type")
	if j != "application/json" {
		t.Fatal(j)
	}
}
