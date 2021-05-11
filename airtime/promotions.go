package airtime

import (
	"encoding/json"
	error2 "github.com/reloadly/reloadly-sdk-golang/error"
	"net/http"
	"strconv"
)

//Promotion is a struct representing a single Promotion.
type Promotion struct {
	ID                 int         `json:"id"`
	PromotionID        int         `json:"promotionId"`
	OperatorID         int         `json:"operatorId"`
	Title              string      `json:"title"`
	Title2             string      `json:"title2"`
	Description        string      `json:"description"`
	StartDate          string      `json:"startDate"`
	EndDate            string      `json:"endDate"`
	Denominations      string      `json:"denominations"`
	LocalDenominations interface{} `json:"localDenominations"`
}

//Promotions is a struct that represents a list of Promotions
type Promotions struct {
	Content []Promotion`json:"content"`
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

//GetPromotions Is used to list all available promotions. Reloady is also supported for enabling operator's promotions. These are provided by the Operators and can be activated by sending a specific topup amount as per the details of the promotion. Reloadly provide neat ways to get all details on the promotion and to showcase these to your customers.
func (c *Client)GetPromotions(filter ...Filters)(*Promotions, error){
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

	requestUrl := c.BaseURL + "/promotions" + query
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.Header.Add("Authorization", c.AuthHeader)
	req.Header.Add("Reloadly-Client", c.telemetryHeader)
	res, err := client.Do(req)

	if err != nil {
		return nil, &error2.ErrorResponse{Message: err.Error()}
	}

	defer res.Body.Close()

	var e error2.ErrorResponse
	var r Promotions
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

//GetPromotionsByID Fetches a Promotion with the specified ID
func (c *Client)GetPromotionsById(Id int)(*Promotion, error){
	method := "GET"
	client := c.HttpClient

	requestUrl := c.BaseURL + "/promotions/" + strconv.Itoa(Id)
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.Header.Add("Authorization", c.AuthHeader)

	res, err := client.Do(req)

	if err != nil {
		return nil, &error2.ErrorResponse{Message: err.Error()}
	}

	defer res.Body.Close()

	var e error2.ErrorResponse
	var r Promotion
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

//GetPromotionsByOperatorId  fetches all the Promotions for the operator id
func (c *Client)GetPromotionsByOperatorId(OperatorId int)(*[]Promotion, error){
	method := "GET"
	client := c.HttpClient

	requestUrl := c.BaseURL + "/promotions/operators/" + strconv.Itoa(OperatorId)
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.Header.Add("Authorization", c.AuthHeader)

	res, err := client.Do(req)

	if err != nil {
		return nil, &error2.ErrorResponse{Message: err.Error()}
	}

	defer res.Body.Close()

	var e error2.ErrorResponse
	var r []Promotion
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

//GetPromotionsByCountryCode fetches all the Promotions for the associated country(code)
func (c *Client)GetPromotionsByCountryCode(CountryCode string)(*[]Promotion, error){
	method := "GET"
	client := c.HttpClient

	requestUrl := c.BaseURL + "/promotions/country-codes/" + CountryCode
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.Header.Add("Authorization", c.AuthHeader)

	res, err := client.Do(req)

	if err != nil {
		return nil, &error2.ErrorResponse{Message: err.Error()}
	}

	defer res.Body.Close()

	var e error2.ErrorResponse
	var r []Promotion
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
