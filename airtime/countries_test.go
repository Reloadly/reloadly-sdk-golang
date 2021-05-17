package airtime_test

import (
	"net/http"
	"testing"
)

func TestClient_GetCountries(t *testing.T) {
		teardown := setup()

		defer teardown()

		mux.HandleFunc("/countries", func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusInternalServerError)
		})

		body, err := client.GetCountries()

		if err == nil {
			t.Errorf("Expected error but got nil")
		}

		if body != nil {
			t.Errorf("Expected body to be nil but got %q",  body)
		}
}

func TestClient_GetCountriesByIso(t *testing.T) {
	teardown := setup()

	defer teardown()

	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
	})

	body, err := client.GetCountriesByIso("")

	if err == nil {
		t.Errorf("Expected error but got nil")
	}

	if body != nil {
		t.Errorf("Expected body to be nil but got %q",  body)
	}
}