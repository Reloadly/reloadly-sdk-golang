package reloadly

import (
	"encoding/json"
	"net/http"
)
type Country struct {
	IsoName        string   `json:"isoName"`
	Name           string   `json:"name"`
	CurrencyCode   string   `json:"currencyCode"`
	CurrencyName   string   `json:"currencyName"`
	CurrencySymbol string   `json:"currencySymbol"`
	Flag           string   `json:"flag"`
	CallingCodes   []string `json:"callingCodes"`
}
type Countries []struct {
	Country Country
}


func (c *Client) GetCountries()(*Countries, error){
	method := "GET"
	client := c.HttpClient

	requestUrl := c.BaseURL + "/countries"
	req, _ := http.NewRequest(method, requestUrl, nil)


	res, err := client.Do(req)

	if err != nil {
		return nil, &Error{Message: err.Error()}
	}

	defer res.Body.Close()

	var e Error
	var r Countries
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

func (c *Client) GetCountriesByIso(ISO string)(*Country, error){
	method := "GET"
	client := c.HttpClient

	requestUrl := c.BaseURL + "/countries/" + ISO
	req, _ := http.NewRequest(method, requestUrl, nil)


	res, err := client.Do(req)

	if err != nil {
		return nil, &Error{Message: err.Error()}
	}

	defer res.Body.Close()

	var e Error
	var r Country
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