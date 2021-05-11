package airtime_test

import (
	"encoding/json"
	reloadly "github.com/reloadly/reloadly-sdk-golang/airtime"
	"net/http"
	"testing"
)

func TestClient_GetDiscounts(t *testing.T) {
	teardown := setup()

	defer teardown()

	mux.HandleFunc("/operators/commissions", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		data := reloadly.Discounts{
			TotalElements: 20,
			TotalPages:    5,
		}
		json.NewEncoder(rw).Encode(data)

	})

	body, err := client.GetDiscounts()
	if err != nil {
		t.Errorf("Expected error to be nil but got %q",  err)
	}

	if body.TotalPages != 5 {
		t.Errorf("Expected TotalPages to be 5 but got %v",  body.TotalPages)
	}

	if body.TotalElements != 20 {
		t.Errorf("Expected TotalElements to be 20 but got %v",  body.TotalElements)
	}
}