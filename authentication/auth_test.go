package authentication_test

import (
	"errors"
	"github.com/Ghvstcode/reloadly/authentication"
	"reflect"
	"testing"
)

type MockAuthClient struct {}
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
			t.Skip("Skipped the rest of the tests")
			if c.ExpectedSandbox != res.SandBox{
				t.Fatalf("Expected Sandbox to be %s but got %s", c.ExpectedSandbox, res.URL)
			}

		}


	}
}
