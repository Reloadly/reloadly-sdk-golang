package reloadly

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type Operators struct {
	Content []struct {
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
			Rate         int    `json:"rate"`
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
	} `json:"content"`
	Pageable struct {
		Sort struct {
			Sorted   bool `json:"sorted"`
			Unsorted bool `json:"unsorted"`
			Empty    bool `json:"empty"`
		} `json:"sort"`
		PageNumber int  `json:"pageNumber"`
		PageSize   int  `json:"pageSize"`
		Offset     int  `json:"offset"`
		Unpaged    bool `json:"unpaged"`
		Paged      bool `json:"paged"`
	} `json:"pageable"`
	TotalElements int  `json:"totalElements"`
	TotalPages    int  `json:"totalPages"`
	Last          bool `json:"last"`
	Sort          struct {
		Sorted   bool `json:"sorted"`
		Unsorted bool `json:"unsorted"`
		Empty    bool `json:"empty"`
	} `json:"sort"`
	First            bool `json:"first"`
	NumberOfElements int  `json:"numberOfElements"`
	Size             int  `json:"size"`
	Number           int  `json:"number"`
	Empty            bool `json:"empty"`
}

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

type OperatorOptions func(opts *OperatorOpts)

func AddPageSize(PageSize int) OperatorOptions {
	return func(s *OperatorOpts) {
		s.PageSize = PageSize
	}
}

func AddPageNumber(PageNumber int) OperatorOptions {
	return func(s *OperatorOpts) {
		s.PageNumber = PageNumber
	}
}

func AddSuggestedAmounts(suggestedAmounts bool) OperatorOptions {
	return func(s *OperatorOpts) {
		s.SuggestedAmounts = suggestedAmounts
	}
}

func AddSuggestedAmountsMap(suggestedAmountsMap bool) OperatorOptions {
	return func(s *OperatorOpts) {
		s.SuggestedAmountsMap = suggestedAmountsMap
	}
}

func AddSimplified(simplified bool) OperatorOptions {
	return func(s *OperatorOpts) {
		s.Simplified = simplified
	}
}

func AddPin(includePin bool) OperatorOptions {
	return func(s *OperatorOpts) {
		s.IncludePin = includePin
	}
}

func AddData(includeData bool) OperatorOptions {
	return func(s *OperatorOpts) {
		s.IncludeData = includeData
	}
}

func AddBundles(includeBundles bool) OperatorOptions {
	return func(s *OperatorOpts) {
		s.IncludeBundles = includeBundles
	}
}

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

	res, err := client.Do(req)

	if err != nil {
		return nil, &Error{Message: err.Error()}
	}

	defer res.Body.Close()

	var e Error
	var r Operators
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &Error{Message: err.Error()}
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &Error{Message: err.Error()}
	}

	return &r, nil
}

func (c *Client) GetOperatorsById(operatorID int, options ...OperatorOptions)(*Operators,error){
	o := &OperatorOpts{}
	for _, opt := range options {
		opt(o)
	}


	method := "GET"
	client := c.HttpClient
	query := fmt.Sprintf("?suggestedAmounts=%t&suggestedAmountsMap=%t", o.SuggestedAmounts,o.SuggestedAmountsMap)
	requestUrl := c.BaseURL + "/operators" + strconv.Itoa(operatorID) + query
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.Header.Add("Authorization", c.AuthHeader)

	res, err := client.Do(req)

	if err != nil {
		return nil, &Error{Message: err.Error()}
	}

	defer res.Body.Close()

	var e Error
	var r Operators
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &Error{Message: err.Error()}
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &Error{Message: err.Error()}
	}

	return &r, nil
}

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
		return nil, &Error{Message: err.Error()}
	}

	defer res.Body.Close()

	var e Error
	var r Operators
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &Error{Message: err.Error()}
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &Error{Message: err.Error()}
	}

	return &r, nil
}

func (c *Client) GetOperatorsByPhone(phone, countryCode string, options ...OperatorOptions)(*Operators,error){
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
		return nil, &Error{Message: err.Error()}
	}

	defer res.Body.Close()

	var e Error
	var r Operators
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &Error{Message: err.Error()}
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &Error{Message: err.Error()}
	}

	return &r, nil
}