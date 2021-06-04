package airtime

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	error2 "github.com/reloadly/reloadly-sdk-golang/error"
	"net/http"
	"strconv"
)

//Operator represents a single Operator
type Operator struct {
	ID                        int         `json:"id"`
	OperatorID                int         `json:"operatorId"`
	Name                      string      `json:"name"`
	Bundle                    bool        `json:"bundle"`
	Data                      bool        `json:"data"`
	Pin                       bool        `json:"pin"`
	SupportsLocalAmounts      bool        `json:"supportsLocalAmounts"`
	DenominationType          string      `json:"denominationType"`
	SenderCurrencyCode        string      `json:"senderCurrencyCode"`
	SenderCurrencySymbol      string      `json:"senderCurrencySymbol"`
	DestinationCurrencyCode   string      `json:"destinationCurrencyCode"`
	DestinationCurrencySymbol string      `json:"destinationCurrencySymbol"`
	Commission                float64     `json:"commission"`
	InternationalDiscount     float64     `json:"internationalDiscount"`
	LocalDiscount             float64     `json:"localDiscount"`
	MostPopularAmount         float64     `json:"mostPopularAmount"`
	MostPopularLocalAmount    interface{} `json:"mostPopularLocalAmount"`
	MinAmount                 interface{} `json:"minAmount"`
	MaxAmount                 interface{} `json:"maxAmount"`
	LocalMinAmount            interface{} `json:"localMinAmount"`
	LocalMaxAmount            interface{} `json:"localMaxAmount"`
	Country                   struct {
		IsoName string `json:"isoName"`
		Name    string `json:"name"`
	} `json:"country"`
	Fx struct {
		Rate         interface{}    `json:"rate"`
		CurrencyCode string `json:"currencyCode"`
	} `json:"fx"`
	LogoUrls                 []string  `json:"logoUrls"`
	FixedAmounts             []float64 `json:"fixedAmounts"`
	FixedAmountsDescriptions struct {
	} `json:"fixedAmountsDescriptions"`
	LocalFixedAmounts             []interface{} `json:"localFixedAmounts"`
	LocalFixedAmountsDescriptions struct {
	} `json:"localFixedAmountsDescriptions"`
	SuggestedAmounts    []interface{} `json:"suggestedAmounts"`
	SuggestedAmountsMap struct {
	} `json:"suggestedAmountsMap"`
	Promotions []interface{} `json:"promotions"`
}

//Operators struct represents list of Operators returned by Relaodly
type Operators struct {
	ID                                int         `json:"id"`
	OperatorID                        int         `json:"operatorId"`
	Name                              string      `json:"name"`
	Bundle                            bool        `json:"bundle"`
	Data                              bool        `json:"data"`
	Pin                               bool        `json:"pin"`
	SupportsLocalAmounts              bool        `json:"supportsLocalAmounts"`
	SupportsGeographicalRechargePlans bool        `json:"supportsGeographicalRechargePlans"`
	DenominationType                  string      `json:"denominationType"`
	SenderCurrencyCode                string      `json:"senderCurrencyCode"`
	SenderCurrencySymbol              string      `json:"senderCurrencySymbol"`
	DestinationCurrencyCode           string      `json:"destinationCurrencyCode"`
	DestinationCurrencySymbol         string      `json:"destinationCurrencySymbol"`
	Commission                        float64     `json:"commission"`
	InternationalDiscount             float64     `json:"internationalDiscount"`
	LocalDiscount                     float64     `json:"localDiscount"`
	MostPopularAmount                 interface{} `json:"mostPopularAmount"`
	MostPopularLocalAmount            interface{} `json:"mostPopularLocalAmount"`
	MinAmount                         interface{} `json:"minAmount"`
	MaxAmount                         interface{} `json:"maxAmount"`
	LocalMinAmount                    interface{} `json:"localMinAmount"`
	LocalMaxAmount                    interface{} `json:"localMaxAmount"`
	Country                           struct {
		IsoName string `json:"isoName"`
		Name    string `json:"name"`
	} `json:"country"`
	Fx struct {
		Rate         int    `json:"rate"`
		CurrencyCode string `json:"currencyCode"`
	} `json:"fx"`
	LogoUrls                 []string      `json:"logoUrls"`
	FixedAmounts             []interface{} `json:"fixedAmounts"`
	FixedAmountsDescriptions struct {
	} `json:"fixedAmountsDescriptions"`
	LocalFixedAmounts             []interface{} `json:"localFixedAmounts"`
	LocalFixedAmountsDescriptions struct {
	} `json:"localFixedAmountsDescriptions"`
	SuggestedAmounts    []interface{} `json:"suggestedAmounts"`
	SuggestedAmountsMap struct {
	} `json:"suggestedAmountsMap"`
	GeographicalRechargePlans []struct {
		LocationCode             string    `json:"locationCode"`
		LocationName             string    `json:"locationName"`
		FixedAmounts             []float64 `json:"fixedAmounts"`
		LocalAmounts             []float64 `json:"localAmounts"`
		FixedAmountsDescriptions struct {} `json:"fixedAmountsDescriptions"`
		LocalFixedAmountsDescriptions struct {} `json:"localFixedAmountsDescriptions"`
	} `json:"geographicalRechargePlans"`
	Promotions []interface{} `json:"promotions"`
}



//OperatorFXRate represents FXRate returned by the Reloadly API
type OperatorFXRate struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	FxRate       int    `json:"fxRate"`
	CurrencyCode string `json:"currencyCode"`
}

//OperatorOpts struct represents the various possible optional parameters
type OperatorOpts struct{
	PageSize int
	PageNumber int
	SuggestedAmounts bool
	SuggestedAmountsMap bool
	Simplified bool
	IncludePin bool
	IncludeData bool
	IncludeBundles bool
}

//OperatorOptions is the Operator Option function type
type OperatorOptions func(opts *OperatorOpts)

//AddPageSize - Specify Page Size
func AddPageSize(PageSize int) OperatorOptions {
	return func(s *OperatorOpts) {
		s.PageSize = PageSize
	}
}

//AddPageNumber - Specify PAge Number
func AddPageNumber(PageNumber int) OperatorOptions {
	return func(s *OperatorOpts) {
		s.PageNumber = PageNumber
	}
}

//AddSuggestedAmounts - Whether to return the suggestedAmounts field on the Operator resource.
func AddSuggestedAmounts(suggestedAmounts bool) OperatorOptions {
	return func(s *OperatorOpts) {
		s.SuggestedAmounts = suggestedAmounts
	}
}

//AddSuggestedAmountsMap - Whether to return the suggestedAmountsMap field on the Operator resource.
func AddSuggestedAmountsMap(suggestedAmountsMap bool) OperatorOptions {
	return func(s *OperatorOpts) {
		s.SuggestedAmountsMap = suggestedAmountsMap
	}
}

//AddSimplified - Whether to return Simplified response or Detailed one.
func AddSimplified(simplified bool) OperatorOptions {
	return func(s *OperatorOpts) {
		s.Simplified = simplified
	}
}

//AddPin - Whether to include PIN details in the operators resources list.
func AddPin(includePin bool) OperatorOptions {
	return func(s *OperatorOpts) {
		s.IncludePin = includePin
	}
}

//AddData - Whether to include airtime/data bundles in the operators resources list.
func AddData(includeData bool) OperatorOptions {
	return func(s *OperatorOpts) {
		s.IncludeData = includeData
	}
}

//AddBundles - Whether to include airtime/data bundles in the operators resources list.
func AddBundles(includeBundles bool) OperatorOptions {
	return func(s *OperatorOpts) {
		s.IncludeBundles = includeBundles
	}
}

//GetOperators retrieves a complete list of all operators.You can set the number of operators to retrieve in each page by simply tweaking the size parameter. Reloadly returns complete detail of each operator so you don't need to make multiple calls, Including what type of operator this is, what topup types it support and even details on the commissions for the operator.
//
//Within the reloadly platform, There exist two types of Operators. One that support Range values (Anything between the minnimun and maximum range). While the other that support Fixed values (Only a cetain values are supported). Reloadly will return you the type of the operator within the response in denominationType variable. If this is set to RANGE you will receive the minimum and maximum values in the minAmount and maxAmount variables for that operator. However if the denomination type is FIXED you will not get these values but rather get an array of all values supported in the fixedAmounts variable. Now a point to remember here is that these values are already converted into your account's currency.
func (c *Client) GetOperators(options ...OperatorOptions)(*Operators,error){
	o := &OperatorOpts{}
	for _, opt := range options {
		opt(o)
	}

	var (
		page string
		size string
	)

	if o.PageNumber == 0 {
		page = "1"
	} else {
		page = strconv.Itoa(o.PageNumber)
	}

	if o.PageSize == 0 {
		size = "200"
	} else{
		size = strconv.Itoa(o.PageSize)
	}

	method := "GET"
	client := c.HttpClient
	query := fmt.Sprintf("?page=%v&size=%v&includeData=%t&includePin=%t&simplified=%t&suggestedAmounts=%t&suggestedAmountsMap=%t&includeBundle=%t", page,size, !o.IncludeData, !o.IncludePin, o.Simplified, o.SuggestedAmounts,o.SuggestedAmountsMap,!o.IncludeBundles)
	requestUrl := c.BaseURL + "/operators" + query
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.Header.Add("Authorization", c.AuthHeader)
	req.Header.Add("Reloadly-Client", c.telemetryHeader)
	res, err := client.Do(req)

	if err != nil {
		return nil, &error2.ErrorResponse{Message: err.Error()}
	}

	defer res.Body.Close()

	var e error2.ErrorResponse
	var r Operators
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &error2.ErrorResponse{Message: err.Error()}
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &error2.ErrorResponse{Message: err.Error()}
	}

	return &r, nil
}

//GetOperatorsById retrieves a specific Operator with the specified ID
func (c *Client) GetOperatorsById(operatorID int, options ...OperatorOptions)(*Operators,error){
	o := &OperatorOpts{}
	for _, opt := range options {
		opt(o)
	}


	method := "GET"
	client := c.HttpClient
	query := fmt.Sprintf("?suggestedAmounts=%t&suggestedAmountsMap=%t", o.SuggestedAmounts,o.SuggestedAmountsMap)
	requestUrl := c.BaseURL + "/operators/" + strconv.Itoa(operatorID) + query
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.Header.Add("Authorization", c.AuthHeader)

	res, err := client.Do(req)

	if err != nil {
		return nil, &error2.ErrorResponse{Message: err.Error()}
	}

	defer res.Body.Close()

	var e error2.ErrorResponse
	var r Operators
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &error2.ErrorResponse{Message: err.Error()}
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &error2.ErrorResponse{Message: err.Error()}
	}

	return &r, nil
}

//GetOperatorsByISO Retrieves a specified Operator by ISO code
func (c *Client) GetOperatorsByISO(ISO string, options ...OperatorOptions)(*Operators,error){
	o := &OperatorOpts{}
	for _, opt := range options {
		opt(o)
	}

	var (
		page string
		size string
	)

	if o.PageNumber == 0 {
		page = "1"
	} else {
		page = strconv.Itoa(o.PageNumber)
	}

	if o.PageSize == 0 {
		size = "200"
	} else{
		size = strconv.Itoa(o.PageSize)
	}

	method := "GET"
	client := c.HttpClient
	query := fmt.Sprintf("?page=%v&size=%v&includeData=%t&includePin=%t&simplified=%t&suggestedAmounts=%t&suggestedAmountsMap=%t&includeBundle=%t", page,size, !o.IncludeData, !o.IncludePin, o.Simplified, o.SuggestedAmounts,o.SuggestedAmountsMap,!o.IncludeBundles)
	requestUrl := c.BaseURL + "/operators/countries/"+ ISO + query
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.Header.Add("Authorization", c.AuthHeader)

	res, err := client.Do(req)

	if err != nil {
		return nil, &error2.ErrorResponse{Message: err.Error()}
	}

	defer res.Body.Close()

	var e error2.ErrorResponse
	var r Operators
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &error2.ErrorResponse{Message: err.Error()}
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &error2.ErrorResponse{Message: err.Error()}
	}

	return &r, nil
}

//GetOperatorsByPhone is the Auto-Detect Option! Reloadly also provide a simple way for its API's users to automatically detect the operator for any given number. This can be done by making a simple GET request to the /operators/auto-detect/phone/{phone}/countries/{iso} Endpoint. We will need to append the phone number with or without the country code and the Country's ISO in the path and the Reloadly platform will automatically find the operator for this number and send you complete details of that operator. Read the Operator's Endpoint Documentation for more details on this.
func (c *Client) GetOperatorsByPhone(phone, countryCode string, options ...OperatorOptions)(*Operator,error){
	o := &OperatorOpts{}
	for _, opt := range options {
		opt(o)
	}


	if phone == "" || countryCode =="" {
		return nil, errors.New("INVALID_CREDENTIALS")
	}

	method := "GET"
	client := c.HttpClient
	query := fmt.Sprintf("/operators/auto-detect/phone/%v/countries/%v?suggestedAmounts=%t&suggestedAmountsMap=%t", phone,countryCode, o.SuggestedAmounts,o.SuggestedAmountsMap)
	requestUrl := c.BaseURL + query
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.Header.Add("Authorization", c.AuthHeader)

	res, err := client.Do(req)

	if err != nil {
		return nil, &error2.ErrorResponse{Message: err.Error()}
	}

	defer res.Body.Close()

	var e error2.ErrorResponse
	var r Operator
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &error2.ErrorResponse{Message: err.Error()}
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &error2.ErrorResponse{Message: err.Error()}
	}

	return &r, nil
}

//GetFXRate is used to calculate the right amount to send and to estimate what amount will be received on the reciever end!So as an example, If you account is in US Dollar and you are trying to send to a nigerian Operator you can quickly make a Post call to the endpoint and send the operator's Id with the Amount in Canadian Dollar to calculate what amount you will receive in Nigerian Naira
func (c *Client) GetFXRate(operatorId, amount int)(*OperatorFXRate,error){
	method := "POST"

	p := map[string]string{
		"operatorId": strconv.Itoa(operatorId),
		"amount": strconv.Itoa(amount),
	}

	data, _ := json.Marshal(p)

	requestUrl := c.BaseURL + "/operators/fx-rate"

	client := c.HttpClient
	req, _ := http.NewRequest(method, requestUrl, bytes.NewBuffer(data))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", c.AuthHeader)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var e error2.ErrorResponse
	var r OperatorFXRate
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, err
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}