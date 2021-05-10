package reloadly_test

import (
	"encoding/json"
	"github.com/reloadly/reloadly-sdk-golang"
	Err "github.com/reloadly/reloadly-sdk-golang/error"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


var (
	mux    *http.ServeMux
	server *httptest.Server
	client *reloadly.Client
)

func setup()func(){
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = &reloadly.Client{
		HttpClient: server.Client(),
		BaseURL: server.URL,
	}

	return func() {
		server.Close()
	}
}

func TestClient_Topup(t *testing.T) {
	teardown := setup()

	defer teardown()

	mux.HandleFunc("/topups", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		data := reloadly.Transaction{}
		json.NewEncoder(rw).Encode(data)
	})

	_, err := client.Topup("100", "test1", false, reloadly.Phone{})
	if err != nil {
		t.Fatal(err)
	}

}

func TestClient_NautaCubaTopup(t *testing.T) {
	teardown := setup()

	defer teardown()

	mux.HandleFunc("/topups", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
		u := &reloadly.Topuprequest{}
		json.NewDecoder(req.Body).Decode(u)
		data := Err.ErrorResponse{}
		if !strings.Contains(u.RecipientEmail, "@nauta"){
			data = Err.ErrorResponse{
				Message:   "invalid Recipient Email",
			}
		}
		json.NewEncoder(rw).Encode(data)

	})

	_, err := client.NautaCubaTopup("100", "test1", false, "reloadly@gmail.com")


	expectedErr := "invalid Recipient Email"
	if err.Error() != expectedErr {
		t.Errorf("Expected error to be %q but it was %q", expectedErr, err)
	}
}

