package reloadly_test

import (
	"encoding/json"
	"github.com/reloadly/reloadly-sdk-golang"
	"net/http"
	"testing"
)

func TestClient_GetTransactions(t *testing.T) {
	teardown := setup()

	defer teardown()

	mux.HandleFunc("/topups/reports/transactions", func(rw http.ResponseWriter, req *http.Request) {

		rw.WriteHeader(http.StatusOK)
		data := reloadly.Transactions{}
		json.NewEncoder(rw).Encode(data)

	})

	_, err := client.GetTransactions()
	if err != nil {
		t.Errorf("Expected error to be nil but got %q",  err)
	}
}

func TestClient_GetTransactionsByID(t *testing.T) {
	teardown := setup()

	defer teardown()

	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		data := reloadly.Transaction{}
		json.NewEncoder(rw).Encode(data)

	})

	_, err := client.GetTransactionsByID(74)
	if err != nil {
		t.Errorf("Expected error to be nil but got %q",  err)
	}
}
