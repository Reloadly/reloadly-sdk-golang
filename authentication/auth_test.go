package authentication_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/reloadly/reloadly-sdk-golang/authentication"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *authentication.AuthClient
)


func setup() func() {
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
	cases := []struct {
		clientId        string
		clientSecret    string
		sandBox         bool
		ExpectedErr     error
		ExpectedSandbox bool
	}{
		{
			clientId:        "1234",
			clientSecret:    "",
			sandBox:         false,
			ExpectedErr:     errors.New("INVALID_CREDENTIALS"),
			ExpectedSandbox: false,
		},
		{
			clientId:        "1234",
			clientSecret:    "5678",
			sandBox:         true,
			ExpectedErr:     nil,
			ExpectedSandbox: true,
		},
	}

	for _, c := range cases {
		res, err := authentication.NewAuthClient(c.clientId, c.clientSecret, c.sandBox)
		if !reflect.DeepEqual(err, c.ExpectedErr) {
			t.Fatalf("Expected err to be %q but it was %q", c.ExpectedErr, err)
		}

		if res != nil {
			//t.Skip("Skipped the rest of the tests")
			if c.ExpectedSandbox != res.SandBox {
				t.Fatalf("Expected Sandbox to be %t but got %s", c.ExpectedSandbox, res.URL)
			}

		}

	}
}


func TestAuthClient_ConfigureHTTP(t *testing.T) {
	type fields struct {
		HttpClient   *http.Client
	}
	type args struct {
		h *http.Client
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Configure HTTP",
			fields: fields{
				HttpClient: http.DefaultClient,
			},
			args: args{
				h: http.DefaultClient,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authentication.AuthClient{
				HttpClient:   tt.fields.HttpClient,
			}
			a.ConfigureHTTP(tt.args.h)
		})
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