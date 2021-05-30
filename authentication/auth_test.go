package authentication_test

import (
	"errors"
	"github.com/reloadly/reloadly-sdk-golang/authentication"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *authentication.AuthClient
)

func setup()func(){
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = &authentication.AuthClient{
		HttpClient: server.Client(),
		URL: server.URL,
	}

	return func() {
		server.Close()
	}
}

func TestNewAuthClient(t *testing.T) {
	cases := [] struct{
		clientId string
		clientSecret string
		sandBox bool
		ExpectedErr error
		ExpectedSandbox bool
	}{
		{
			clientId: "1234",
			clientSecret: "",
			sandBox: false,
			ExpectedErr: errors.New("INVALID_CREDENTIALS"),
			ExpectedSandbox: false,
		},
		{
			clientId: "1234",
			clientSecret: "5678",
			sandBox: true,
			ExpectedErr: nil,
			ExpectedSandbox: true,
		},
	}

	for _, c := range cases {
		res, err := authentication.NewAuthClient(c.clientId,c.clientSecret, c.sandBox)
		if !reflect.DeepEqual(err, c.ExpectedErr) {
			t.Fatalf("Expected err to be %q but it was %q", c.ExpectedErr, err)
		}

		if res != nil{
			//t.Skip("Skipped the rest of the tests")
			if c.ExpectedSandbox != res.SandBox{
				t.Fatalf("Expected Sandbox to be %t but got %s", c.ExpectedSandbox, res.URL)
			}

		}


	}
}


func TestAuthClient_GetAccessToken(t *testing.T) {
	teardown := setup()

	defer teardown()

	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
	})

	body, err := client.GetAccessToken()

	if err == nil {
		t.Errorf("Expected error but got nil")
	}

	if body != nil {
		t.Errorf("Expected body to be nil but got %v",  body)
	}
}