package airtime

import (
	"encoding/json"
	error2 "github.com/reloadly/reloadly-sdk-golang/error"
	"net/http"
)

//Country struct represents a Country response from the Reloadly API.
type Country struct {
	IsoName        string   `json:"isoName"`
	Name           string   `json:"name"`
	CurrencyCode   string   `json:"currencyCode"`
	CurrencyName   string   `json:"currencyName"`
	CurrencySymbol string   `json:"currencySymbol"`
	Flag           string   `json:"flag"`
	CallingCodes   []string `json:"callingCodes"`
}

//GetCountries is used to retrieve a list of countries the Reloadly API supports
func (c *Client) GetCountries()(*[]Country, error){
	method := "GET"
	client := c.HttpClient

	requestUrl := c.BaseURL + "/countries"
	req, _ := http.NewRequest(method, requestUrl, nil)
	req.Header.Add("Reloadly-Client", c.telemetryHeader)

	res, err := client.Do(req)

	if err != nil {
		return nil, &error2.ErrorResponse{Message: err.Error()}
	}

	defer res.Body.Close()

	var e error2.ErrorResponse
	var r []Country
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

//GetCountriesByIso is used to retreive a specific country with the corresponding IDO code
func (c *Client) GetCountriesByIso(ISO string)(*Country, error){
	method := "GET"
	client := c.HttpClient

	requestUrl := c.BaseURL + "/countries/" + ISO
	req, _ := http.NewRequest(method, requestUrl, nil)
	req.Header.Add("Reloadly-Client", c.telemetryHeader)

	res, err := client.Do(req)

	if err != nil {
		return nil, &error2.ErrorResponse{Message: err.Error()}
	}

	defer res.Body.Close()

	var e error2.ErrorResponse
	var r Country
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