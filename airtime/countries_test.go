package airtime_test

import (
	"encoding/json"
	reloadly "github.com/reloadly/reloadly-sdk-golang/airtime"
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
			t.Errorf("Expected error to be %q but got nil",  err)
		}

		if body != nil {
			t.Errorf("Expected body to be nil but got %q",  body)
		}
}

func TestClient_GetCountriesByIso(t *testing.T) {
	teardown := setup()

	defer teardown()

	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)

		data := reloadly.Country{
			IsoName:        "NG",
			CurrencyCode:   "NGN",
			CurrencyName:   "Naira",
		}

		json.NewEncoder(rw).Encode(data)
	})


	body, err := client.GetCountriesByIso("NG")

	if err != nil {
		t.Errorf("Expected error to be %q but got nil",  err)
	}

	if body.IsoName != "NG" {
		t.Errorf("Expected body to be NG but got %q",  body.IsoName)
	}
}