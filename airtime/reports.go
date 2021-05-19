package airtime

import (
	"encoding/json"
	error2 "github.com/reloadly/reloadly-sdk-golang/error"
	"net/http"
	"strconv"
)

//Transaction represents a single Transaction made!
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
	RequestedAmount             interface{}         `json:"requestedAmount"`
	RequestedAmountCurrencyCode string      `json:"requestedAmountCurrencyCode"`
	DeliveredAmount             float64     `json:"deliveredAmount"`
	DeliveredAmountCurrencyCode string      `json:"deliveredAmountCurrencyCode"`
	TransactionDate             string      `json:"transactionDate"`
	PinDetail                   interface{} `json:"pinDetail"`
	BalanceInfo                 interface{} `json:"balanceInfo"`
}

//Transactions Is the list of all transactions made using Reloadly
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

//GetTransactions is used to retrieve a list of all past transactions.
//Just like all systems, We also keep a record of all transactions so that our users can track their activity in their respective dashboards.
//Apart from doing just that we also provide a neat way to integrate transactions into your own systems.
//You can get a complete list of your past top-up transactions by calling the GetTransactions function. Internally, it makes a GET request to the /topups/reports/transactions Endpoint.
//This will provide you with a paginated result with all your recent transactions and the ability to paginate to further previous transactions.

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
	req.Header.Add("Reloadly-Client", c.telemetryHeader)
	res, err := client.Do(req)

	if err != nil {
		return nil, &error2.ErrorResponse{Message: err.Error()}
	}

	defer res.Body.Close()

	var e error2.ErrorResponse
	var r Transactions
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

//GetTransactionsByID fetches a specific transaction By it's ID
func (c *Client)GetTransactionsByID(Id int)(*Transaction, error){
	method := "GET"
	client := c.HttpClient

	requestUrl := c.BaseURL + "/topups/reports/transactions/" + strconv.Itoa(Id)
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.Header.Add("Authorization", c.AuthHeader)

	res, err := client.Do(req)

	if err != nil {
		return nil, &error2.ErrorResponse{Message: err.Error()}
	}

	defer res.Body.Close()

	var e error2.ErrorResponse
	var r Transaction
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
