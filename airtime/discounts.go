package airtime

import (
	"encoding/json"
	error2 "github.com/reloadly/reloadly-sdk-golang/error"
	"net/http"
)

type Discounts struct {
	Content []Discount `json:"content"`
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

type Discount struct {
	Percentage              float64 `json:"percentage"`
	InternationalPercentage float64 `json:"internationalPercentage"`
	LocalPercentage         float64 `json:"localPercentage"`
	UpdatedAt               string  `json:"updatedAt"`
	Operator                struct {
		ID          int    `json:"id"`
		OperatorID  int    `json:"operatorId"`
		Name        string `json:"name"`
		CountryCode string `json:"countryCode"`
		Data        bool   `json:"data"`
		Bundle      bool   `json:"bundle"`
		Status      bool   `json:"status"`
	} `json:"operator"`
}


//GetDiscounts Retrieves all available discounts
func (c *Client)GetDiscounts(filter ...Filters)(*Discounts, error){
	o := &FilterOptions{}
	for _, opt := range filter {
		opt(o)
	}
	var query string
	if o.Page == "" && o.Size == ""{
		query = ""
	}

	if o.Page != "" && o.Size != ""{
		query = "?page=" + o.Page + "&size=" + o.Size
	}

	if o.Page != "" && o.Size == "" {
		query = "?page=" + o.Page
	}

	if o.Page == "" && o.Size != "" {
		query = "?size=" + o.Size
	}

	method := "GET"
	client := c.HttpClient

	requestUrl := c.BaseURL + "/operators/commissions" + query
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.Header.Add("Authorization", c.AuthHeader)
	req.Header.Add("Reloadly-Client", c.telemetryHeader)
	res, err := client.Do(req)

	if err != nil {
		return nil, &error2.ErrorResponse{Message: err.Error()}
	}

	defer res.Body.Close()

	var e error2.ErrorResponse
	var r Discounts
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

//GetDiscountsByOperatorID is used to retrieve a specific Discount associated with an Operator
func (c *Client)GetDiscountsByOperatorID(OperatorID string)(*Discounts, error){
	method := "GET"
	client := c.HttpClient

	requestUrl := c.BaseURL + "/operators/" + OperatorID + "/commissions"
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.Header.Add("Authorization", c.AuthHeader)

	res, err := client.Do(req)

	if err != nil {
		return nil, &error2.ErrorResponse{Message: err.Error()}
	}

	defer res.Body.Close()

	var e error2.ErrorResponse
	var r Discounts
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




