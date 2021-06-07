package airtime_test

import (
	"encoding/json"
	"errors"
	reloadly "github.com/reloadly/reloadly-sdk-golang/airtime"
	Err "github.com/reloadly/reloadly-sdk-golang/error"
	"net/http"
	"net/http/httptest"
	"reflect"
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

func TestGetPhone(t *testing.T) {
	cases := [] struct{
		Number string
		CountryCode string
		ExpectedNumber string
		ExpectedCountryCode string
		ExpectedErr error
	}{
		{
			Number: "",
			CountryCode: "",
			ExpectedNumber: "",
			ExpectedCountryCode: "",
			ExpectedErr: errors.New("Invalid_Credentials"),
		},
		{
			Number: "5555",
			CountryCode: "4444",
			ExpectedNumber: "5555",
			ExpectedCountryCode: "4444",
			ExpectedErr: nil,
		},
	}


	for _, c := range cases {
		res, err := reloadly.GetPhone(c.Number, c.CountryCode)
		if !reflect.DeepEqual(err, c.ExpectedErr) {
			t.Fatalf("Expected err to be %q but it was %q", c.ExpectedErr, err)
		}

		if res != nil{
			//t.Skip("Skipped the rest of the tests")
			if c.ExpectedNumber != res.Number {
				t.Fatalf("Expected Number to be %s but got %s", c.ExpectedNumber, res.Number)
			}

			if c.CountryCode != res.CountryCode {
				t.Fatalf("Expected Country Code to be %s but got %s", c.ExpectedCountryCode, res.CountryCode)
			}
		}


	}
}

func TestAddCustomIdentifier(t *testing.T) {
	cases := [] struct{
		CustomIdentifier string
		ExpectedCustomIdentifier string
	}{
		{
			CustomIdentifier: "my-reloadly",
			ExpectedCustomIdentifier: "my-reloadly",
		},
	}



	for _, c := range cases {
		res := reloadly.AddCustomIdentifier(c.CustomIdentifier)
		o := &reloadly.TopupOpts{}
		res(o)

		if res != nil{
			if c.ExpectedCustomIdentifier != o.CustomIdentifier {
				t.Fatalf("Expected Custom Identifier to be %s but got %s", c.ExpectedCustomIdentifier, o.CustomIdentifier)
			}

		}


	}
}

func TestAddSenderPhone(t *testing.T) {
	cases := [] struct{
		SenderPhone reloadly.Phone
		ExpectedSenderPhone reloadly.Phone
	}{
		{
			SenderPhone: reloadly.Phone{
				CountryCode: "NG",
				Number:      "000000",
			},
			ExpectedSenderPhone: reloadly.Phone{
				CountryCode: "NG",
				Number:      "000000",
			},
		},
	}



	for _, c := range cases {
		res := reloadly.AddSenderPhone(c.SenderPhone)
		o := &reloadly.TopupOpts{}
		res(o)

		if res != nil{
			if c.ExpectedSenderPhone != *o.SenderPhone {
				t.Fatalf("Unexpected Sender Phone")
			}

		}


	}
}

