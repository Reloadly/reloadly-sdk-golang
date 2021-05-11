package airtime

import (
	"encoding/json"
	"errors"
	error2 "github.com/reloadly/reloadly-sdk-golang/error"
	"net/http"
	"strings"
)

//Phone is The recipient's Phone object. It is required for topups
type Phone struct {
	CountryCode string `json:"countryCode"`
	Number string `json:"number"`
}

type Topuprequest struct {
	OperatorID       string `json:"operatorId"`
	Amount           string `json:"amount"`
	UseLocalAmount   bool `json:"useLocalAmount"`
	CustomIdentifier string `json:"customIdentifier,omitempty"`
	RecipientPhone   Phone `json:"recipientPhone"`
	RecipientEmail   string `json:"recipientEmail,omitempty"`
	SenderPhone *Phone `json:"senderPhone,omitempty"`
}

//TopupOpts Option struct
type TopupOpts struct{
	CustomIdentifier string
	SenderPhone *Phone
}

//TopupOptions function type
type TopupOptions func(opts *TopupOpts)

func GetPhone(Number, CountryCode string)(*Phone, error){
	if Number =="" || CountryCode == ""{
		return nil, errors.New("Invalid_Credentials")
	}

	return &Phone{
		Number:      Number,
		CountryCode: CountryCode,
	}, nil
}

//AddCustomIdentifier Option is used to add a Custom Identifier for the transaction.
func AddCustomIdentifier(CustomIdentifier string) TopupOptions {
	return func(s *TopupOpts) {
		s.CustomIdentifier = CustomIdentifier
	}
}

//AddSenderPhone Option is used to add a sender's Phone Number when sending a topup.
func AddSenderPhone(SenderPhone Phone) TopupOptions {
	return func(s *TopupOpts) {
		s.SenderPhone = &SenderPhone
	}
}

//Topup is used to TopUP the provided PhoneNumber.
//In order to send a successful topup. There are a few prerequisites to the system. We need to know the phone number to send the topup to, the operator id of the phone number, the country of the operator, the amount for the topup.
func (c *Client) Topup(amount, operatorID string, useLocalAmount bool, recipientPhone Phone, options ...TopupOptions)(*Transaction, error){
	o := &TopupOpts{}
	for _, opt := range options {
		opt(o)
	}

	requestBody := &Topuprequest{
		OperatorID:       operatorID,
		Amount:           amount,
		UseLocalAmount:   useLocalAmount,
		CustomIdentifier: o.CustomIdentifier,
		RecipientPhone:   recipientPhone,
		SenderPhone:      o.SenderPhone,
	}

	requestUrl := c.BaseURL + "/topups"

	rb, err := json.Marshal(requestBody)

	if err != nil {
		return nil, err
	}

	payload := strings.NewReader(string(rb))
	method := "POST"
	client := c.HttpClient
	req, _ := http.NewRequest(method, requestUrl, payload)

	req.Header.Add("Authorization", c.AuthHeader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Reloadly-Client", c.telemetryHeader)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()



	var e error2.ErrorResponse
	var ar Transaction
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, err
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

//NautaCubaTopup Is Reloadly's implementation of Nauta Cuba for top-ups.
//However the process is a bit different from sending phone topups. Instead of using recipientPhone use recipientEmail.
//The rest of the process is exactly the same as sending any other topup.
//Note, There are two types of email domains that are allowed for Nauta Cuba Top-up : @nauta.com.cu and @nauta.co.cu
func (c *Client) NautaCubaTopup(amount, operatorID string, useLocalAmount bool, recipientEmail string, options ...TopupOptions)(*Transaction, error){
	o := &TopupOpts{}
	for _, opt := range options {
		opt(o)
	}

	requestBody := &Topuprequest{
		OperatorID:       operatorID,
		Amount:           amount,
		UseLocalAmount:   useLocalAmount,
		CustomIdentifier: o.CustomIdentifier,
		RecipientEmail:   recipientEmail,
		SenderPhone:      o.SenderPhone,
	}


	rb, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	requestUrl := c.BaseURL + "/topups"
	payload := strings.NewReader(string(rb))
	method := "POST"
	client := c.HttpClient
	req, _ := http.NewRequest(method, requestUrl, payload)

	req.Header.Add("Authorization", c.AuthHeader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Reloadly-Client", c.telemetryHeader)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var e error2.ErrorResponse
	var ar Transaction
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, err
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

