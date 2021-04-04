package reloadly

import (
	"encoding/json"
	error2 "github.com/Ghvstcode/reloadly/error"
	"net/http"
)

type AccountBalance struct {
	Balance      float64 `json:"balance"`
	CurrencyCode string  `json:"currencyCode"`
	CurrencyName string  `json:"currencyName"`
	UpdatedAt    string  `json:"updatedAt"`
}

func (c *Client) GetBalance()(*AccountBalance, error){
	method := "GET"
	client := c.HttpClient

	requestUrl := c.BaseURL + "/accounts/balance"
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.Header.Add("Authorization", c.AuthHeader)

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

