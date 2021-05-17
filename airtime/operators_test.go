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
			Size: 20,
		}

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

func TestClient_GetOperatorsById(t *testing.T) {
	teardown := setup()

	defer teardown()

	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		data := reloadly.Operators{
			TotalPages:    5,
			Size: 20,
		}

		json.NewEncoder(rw).Encode(data)

	})

	body, err := client.GetOperatorsById(7)
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
		rw.WriteHeader(http.StatusInternalServerError)
	})

	body, err := client.GetOperators()

	if err == nil {
		t.Errorf("Expected error but got nil")
	}

	if body != nil {
		t.Errorf("Expected body to be nil but got %v",  body)
	}
}

func TestAddSuggestedAmounts(t *testing.T) {
	cases := [] struct{
		SuggestedAmounts bool
		ExpectedSuggestedAmounts bool
	}{
		{
			SuggestedAmounts: false,
			ExpectedSuggestedAmounts: false,
		},
		{
			SuggestedAmounts: true,
			ExpectedSuggestedAmounts: true,
		},
	}



	for _, c := range cases {
		res := reloadly.AddSuggestedAmounts(c.SuggestedAmounts)
		o := &reloadly.OperatorOpts{}
		res(o)

		if res != nil{
			if c.ExpectedSuggestedAmounts != o.SuggestedAmounts {
				t.Fatalf("Expected SuggestedAmounts to be %t but got %t", c.SuggestedAmounts, o.SuggestedAmounts)
			}

		}


	}
}

func TestAddBundles(t *testing.T) {
	cases := [] struct{
		IncludeBundles bool
		ExpectedIncludeBundles bool
	}{
		{
			IncludeBundles: false,
			ExpectedIncludeBundles: false,
		},
		{
			IncludeBundles: true,
			ExpectedIncludeBundles: true,
		},
	}



	for _, c := range cases {
		res := reloadly.AddBundles(c.IncludeBundles)
		o := &reloadly.OperatorOpts{}
		res(o)

		if res != nil{
			if c.ExpectedIncludeBundles != o.IncludeBundles {
				t.Fatalf("Expected SuggestedAmounts to be %t but got %t", c.IncludeBundles, o.IncludeBundles)
			}

		}


	}
}

func TestAddData(t *testing.T) {
	cases := [] struct{
		IncludeData bool
		ExpectedIncludeData bool
	}{
		{
			IncludeData: false,
			ExpectedIncludeData: false,
		},
		{
			IncludeData: true,
			ExpectedIncludeData: true,
		},
	}



	for _, c := range cases {
		res := reloadly.AddData(c.IncludeData)
		o := &reloadly.OperatorOpts{}
		res(o)

		if res != nil{
			if c.IncludeData != o.IncludeData {
				t.Fatalf("Expected SuggestedAmounts to be %t but got %t", c.IncludeData, o.IncludeData)
			}

		}


	}
}