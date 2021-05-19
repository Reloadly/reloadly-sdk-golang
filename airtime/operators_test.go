package airtime_test

import (
	"encoding/json"
	reloadly "github.com/reloadly/reloadly-sdk-golang/airtime"
	"net/http"
	"testing"
)

func TestClient_GetOperatorsByPhone(t *testing.T) {
	teardown := setup()

	defer teardown()

	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		data := reloadly.Operator{
			OperatorID: 321,
		}
		json.NewEncoder(rw).Encode(data)

	})
	body, err := client.GetOperatorsByPhone("+123456", "NG")
	if err != nil {
		t.Errorf("Expected error to be nil but got %q",  err)
	}

	if body.OperatorID != 321 {
		t.Errorf("Expected OperatorID to be 321 but got %v",  body.OperatorID)
	}
}

func TestClient_GetFXRate(t *testing.T) {
	teardown := setup()

	defer teardown()

	mux.HandleFunc("/operators/fx-rate", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		data := reloadly.OperatorFXRate{
			ID:           1,
			Name:         "Afghan Wireless",
			FxRate:       65,
			CurrencyCode: "AFN",
		}
		json.NewEncoder(rw).Encode(data)

	})

	body, err := client.GetFXRate(1, 1)
	if err != nil {
		t.Errorf("Expected error to be nil but got %q",  err)
	}

	if body.FxRate != 65 {
		t.Errorf("Expected OperatorID to be 65 but got %v",  body.FxRate)
	}
}

func TestClient_GetOperatorsByISO(t *testing.T) {
	teardown := setup()

	defer teardown()

	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		data := reloadly.Operators{
			TotalPages:    5,
			Size:             20,
		}
		//data := reloadly.Operators{
		//	TotalPages:    5,
		//	Size: 20,
		//}

		json.NewEncoder(rw).Encode(data)

	})

	body, err := client.GetOperatorsByISO("011")
	if err != nil {
		t.Errorf("Expected error to be nil but got %q",  err)
	}

	if body.Size != 20 {
		t.Errorf("Expected Size to be 20 but got %v",  body.Size)
	}
}

func TestClient_GetOperators(t *testing.T) {
	teardown := setup()

	defer teardown()

	mux.HandleFunc("/operators", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		data := reloadly.Operators{
			TotalPages:    5,
			Size: 20,
		}

		json.NewEncoder(rw).Encode(data)

	})

	body, err := client.GetOperators()
	if err != nil {
		t.Errorf("Expected error to be nil but got %q",  err)
	}

	if body.Size != 20 {
		t.Errorf("Expected Size to be 20 but got %v",  body.Size)
	}
}

func TestClient_GetOperatorsById(t *testing.T) {
	teardown := setup()

	defer teardown()

	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)

	})

	_, err := client.GetOperators()
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}