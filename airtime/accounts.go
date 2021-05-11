package airtime

import (
	"encoding/json"
	error2 "github.com/reloadly/reloadly-sdk-golang/error"
	"net/http"
)

//AccountBalance struct represents the account balance response.
type AccountBalance struct {
	Balance      float64 `json:"balance"`
	CurrencyCode string  `json:"currencyCode"`
	CurrencyName string  `json:"currencyName"`
	UpdatedAt    string  `json:"updatedAt"`
}

//GetBalance is used to list account available balance
func (c *Client) GetBalance()(*AccountBalance, error){
	method := "GET"
	client := c.HttpClient

	requestUrl := c.BaseURL + "/accounts/balance"
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.Header.Add("Authorization", c.AuthHeader)
	req.Header.Add("Reloadly-Client", c.telemetryHeader)
	res, err := client.Do(req)

	if err != nil {
		return nil, &error2.ErrorResponse{Message: err.Error()}
	}

	defer res.Body.Close()

	var e error2.ErrorResponse
	var r AccountBalance
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

