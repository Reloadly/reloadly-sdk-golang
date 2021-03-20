package reloadly

import (
	"encoding/json"
	"net/http"
	"strconv"
)

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

	res, err := client.Do(req)

	if err != nil {
		return nil, &Error{Message: err.Error()}
	}

	defer res.Body.Close()

	var e Error
	var r Promotions
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

func (c *Client)GetPromotionsById(Id int)(*Promotion, error){
	method := "GET"
	client := c.HttpClient

	requestUrl := c.BaseURL + "/promotions" + strconv.Itoa(Id)
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.Header.Add("Authorization", c.AuthHeader)

	res, err := client.Do(req)

	if err != nil {
		return nil, &Error{Message: err.Error()}
	}

	defer res.Body.Close()

	var e Error
	var r Promotion
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

func (c *Client)GetPromotionsByOperatorId(OperatorId int)(*Promotion, error){
	method := "GET"
	client := c.HttpClient

	requestUrl := c.BaseURL + "/promotions" + strconv.Itoa(OperatorId)
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.Header.Add("Authorization", c.AuthHeader)

	res, err := client.Do(req)

	if err != nil {
		return nil, &Error{Message: err.Error()}
	}

	defer res.Body.Close()

	var e Error
	var r Promotion
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

func (c *Client)GetPromotionsByCode(CountryCode string)(*Promotion, error){
	method := "GET"
	client := c.HttpClient

	requestUrl := c.BaseURL + "/promotions" + CountryCode
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.Header.Add("Authorization", c.AuthHeader)

	res, err := client.Do(req)

	if err != nil {
		return nil, &Error{Message: err.Error()}
	}

	defer res.Body.Close()

	var e Error
	var r Promotion
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
