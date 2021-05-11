package airtime_test

import (
	"encoding/json"
	reloadly "github.com/reloadly/reloadly-sdk-golang/airtime"
	"net/http"
	"testing"
)

func TestClient_GetBalance(t *testing.T) {
	teardown := setup()

	defer teardown()

	mux.HandleFunc("/accounts/balance", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		data :=  reloadly.AccountBalance{Balance: 950}
		json.NewEncoder(rw).Encode(data)

	})

	body, err := client.GetBalance()
	if err != nil {
		t.Errorf("Expected error to be nil but got %q",  err)
	}

	if body.Balance != 950 {
		t.Errorf("Expected error to be 950 but got %f",  body.Balance)
	}
}
