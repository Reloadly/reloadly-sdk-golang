package airtime_test

import (
	"encoding/json"
	reloadly "github.com/reloadly/reloadly-sdk-golang/airtime"
	"net/http"
	"testing"
)

func TestClient_GetPromotions(t *testing.T) {
	teardown := setup()

	defer teardown()

	mux.HandleFunc("/promotions", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		data := reloadly.Promotions{}
		json.NewEncoder(rw).Encode(data)

	})

	_, err := client.GetPromotions()
	if err != nil {
		t.Errorf("Expected error to be nil but got %q",  err)
	}
}

func TestClient_GetPromotionsByCode(t *testing.T) {
	teardown := setup()

	defer teardown()

	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		var data []reloadly.Promotion
		json.NewEncoder(rw).Encode(data)

	})

	_, err := client.GetPromotionsByCountryCode("NG")
	if err != nil {
		t.Errorf("Expected error to be nil but got %q",  err)
	}
}

func TestClient_GetPromotionsById(t *testing.T) {
	teardown := setup()

	defer teardown()

	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetPromotionsById(0123)
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

func TestClient_GetPromotionsByOperatorId(t *testing.T) {
	teardown := setup()

	defer teardown()

	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetPromotionsByOperatorId(0123)
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

