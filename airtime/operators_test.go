package airtime_test

import (
	"encoding/json"
	"net/http"
	"testing"

	reloadly "github.com/reloadly/reloadly-sdk-golang/airtime"
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
		rw.WriteHeader(http.StatusInternalServerError)
	})

	body, err := client.GetOperatorsByISO("")

	if err == nil {
		t.Errorf("Expected error but got nil")
	}

	if body != nil {
		t.Errorf("Expected body to be nil but got %v",  body)
	}
}


func TestClient_GetOperatorsById(t *testing.T) {
	teardown := setup()

	defer teardown()

	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		data := reloadly.Operators{
			Name: "testOperator",
		}

		json.NewEncoder(rw).Encode(data)

	})

	body, err := client.GetOperatorsById(7)
	if err != nil {
		t.Errorf("Expected error to be nil but got %q",  err)
	}

	if body.Name != "testOperator" {
		t.Errorf("Expected Size to be testOperator but got %v",  body.Name)
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
				t.Fatalf("Expected Add Bundle Option to be %t but got %t", c.IncludeBundles, o.IncludeBundles)
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
			if c.ExpectedIncludeData != o.IncludeData {
				t.Fatalf("Expected Add Data Option to be %t but got %t", c.ExpectedIncludeData, o.IncludeData)
			}

		}


	}
}

func TestAddPageSize(t *testing.T) {
	cases := [] struct{
		PageSize int
		ExpectedPageSize int
	}{
		{
			PageSize: 4,
			ExpectedPageSize: 4,
		},
	}



	for _, c := range cases {
		res := reloadly.AddPageSize(c.PageSize)
		o := &reloadly.OperatorOpts{}
		res(o)

		if res != nil{
			if c.PageSize != o.PageSize {
				t.Fatalf("Expected SuggestedAmounts to be %d but got %d", c.PageSize, o.PageSize)
			}

		}


	}
}

func TestAddPageNumber(t *testing.T) {
	cases := [] struct{
		PageNumber int
		ExpectedPageNumber int
	}{
		{
			PageNumber: 5,
			ExpectedPageNumber: 5,
		},
	}



	for _, c := range cases {
		res := reloadly.AddPageNumber(c.PageNumber)
		o := &reloadly.OperatorOpts{}
		res(o)

		if res != nil{
			if c.PageNumber != o.PageNumber {
				t.Fatalf("Expected SuggestedAmounts to be %d but got %d", c.PageNumber, o.PageNumber)
			}

		}


	}
}

func TestAddPin(t *testing.T) {
	cases := [] struct{
		IncludePin bool
		ExpectedIncludePin bool
	}{
		{
			IncludePin: false,
			ExpectedIncludePin: false,
		},
	}



	for _, c := range cases {
		res := reloadly.AddPin(c.IncludePin)
		o := &reloadly.OperatorOpts{}
		res(o)

		if res != nil{
			if c.ExpectedIncludePin != o.IncludePin {
				t.Fatalf("Expected Include Pin Option to be %t but got %t", c.ExpectedIncludePin, o.IncludePin)
			}

		}


	}
}

func TestAddSuggestedAmountsMap(t *testing.T) {
	cases := [] struct{
		SuggestedAmountsMap bool
		ExpectedSuggestedAmountsMap bool
	}{
		{
			SuggestedAmountsMap: false,
			ExpectedSuggestedAmountsMap: false,
		},
	}



	for _, c := range cases {
		res := reloadly.AddSuggestedAmountsMap(c.SuggestedAmountsMap)
		o := &reloadly.OperatorOpts{}
		res(o)

		if res != nil{
			if c.SuggestedAmountsMap != o.SuggestedAmountsMap {
				t.Fatalf("Expected Suggested Amounts Map Option to be %t but got %t", c.SuggestedAmountsMap, o.SuggestedAmountsMap)
			}

		}


	}
}

func TestAddSimplified(t *testing.T) {
	cases := [] struct{
		AddSimplified bool
		ExpectedAddSimplified bool
	}{
		{
			AddSimplified: false,
			ExpectedAddSimplified: false,
		},
	}



	for _, c := range cases {
		res := reloadly.AddSimplified(c.AddSimplified)
		o := &reloadly.OperatorOpts{}
		res(o)

		if res != nil{
			if c.ExpectedAddSimplified != o.Simplified {
				t.Fatalf("Expected Simplified Option to be %t but got %t", c.AddSimplified, o.Simplified)
			}

		}


	}
}