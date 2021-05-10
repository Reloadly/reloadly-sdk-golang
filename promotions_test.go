package reloadly_test

import (
	"encoding/json"
	"github.com/reloadly/reloadly"
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

