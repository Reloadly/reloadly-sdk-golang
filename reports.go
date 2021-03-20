package reloadly

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Transaction struct {
	TransactionID               int         `json:"transactionId"`
	OperatorTransactionID       interface{} `json:"operatorTransactionId"`
	CustomIdentifier            string      `json:"customIdentifier"`
	RecipientPhone              string      `json:"recipientPhone"`
	RecipientEmail              interface{} `json:"recipientEmail"`
	SenderPhone                 string      `json:"senderPhone"`
	CountryCode                 string      `json:"countryCode"`
	OperatorID                  int         `json:"operatorId"`
	OperatorName                string      `json:"operatorName"`
	Discount                    float64     `json:"discount"`
	DiscountCurrencyCode        string      `json:"discountCurrencyCode"`
	RequestedAmount             int         `json:"requestedAmount"`
	RequestedAmountCurrencyCode string      `json:"requestedAmountCurrencyCode"`
	DeliveredAmount             float64     `json:"deliveredAmount"`
	DeliveredAmountCurrencyCode string      `json:"deliveredAmountCurrencyCode"`
	TransactionDate             string      `json:"transactionDate"`
	PinDetail                   interface{} `json:"pinDetail"`
	BalanceInfo                 interface{} `json:"balanceInfo"`
}

type Transactions struct {
	Content []Transaction `json:"content"`
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

func (c *Client)GetTransactions(filter ...Filters)(*Transactions, error){
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

	requestUrl := c.BaseURL + "/topups/reports/transactions" + query
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.Header.Add("Authorization", c.AuthHeader)

	res, err := client.Do(req)

	if err != nil {
		return nil, &Error{Message: err.Error()}
	}

	defer res.Body.Close()

	var e Error
	var r Transactions
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

func (c *Client)GetTransactionsByID(Id int)(*Transactions, error){
	method := "GET"
	client := c.HttpClient

	requestUrl := c.BaseURL + "/topups/reports/transactions" + strconv.Itoa(Id)
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.Header.Add("Authorization", c.AuthHeader)

	res, err := client.Do(req)

	if err != nil {
		return nil, &Error{Message: err.Error()}
	}

	defer res.Body.Close()

	var e Error
	var r Transactions
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
