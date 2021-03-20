package reloadly

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type Phone struct {
	Number string
	CountryCode string
}

type request struct {
	OperatorID       string `json:"operatorId"`
	Amount           string `json:"amount"`
	UseLocalAmount   bool `json:"useLocalAmount"`
	CustomIdentifier string `json:"customIdentifier,omitempty"`
	RecipientPhone   Phone `json:"recipientPhone"`
	RecipientEmail   string `json:"recipientEmail,omitempty"`
	SenderPhone *Phone `json:"senderPhone,omitempty"`
}

type TopupOpts struct{
	CustomIdentifier string
	SenderPhone Phone
}

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

func AddCustomIdentifier(CustomIdentifier string) TopupOptions {
	return func(s *TopupOpts) {
		s.CustomIdentifier = CustomIdentifier
	}
}

func AddSenderPhone(SenderPhone Phone) TopupOptions {
	return func(s *TopupOpts) {
		s.SenderPhone = SenderPhone
	}
}

func (c *Client) Topup(amount, operatorID string, useLocalAmount bool, recipientPhone Phone, options ...TopupOptions)(*Transaction, error){
	o := &TopupOpts{}
	for _, opt := range options {
		opt(o)
	}

	requestBody := &request{
		OperatorID:       operatorID,
		Amount:           amount,
		UseLocalAmount:   useLocalAmount,
		CustomIdentifier: o.CustomIdentifier,
		RecipientPhone:   recipientPhone,
		SenderPhone:      &o.SenderPhone,
	}


	rb, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	payload := strings.NewReader(string(rb))
	method := "POST"
	client := c.HttpClient
	req, _ := http.NewRequest(method, c.BaseURL, payload)

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var e error
	var ar Transaction
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, err
		}
		return nil, e

	}

	err = json.NewDecoder(res.Body).Decode(&ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

func (c *Client) NautaCubaTopup(amount, operatorID string, useLocalAmount bool, recipientEmail string, options ...TopupOptions)(*Transaction, error){
	o := &TopupOpts{}
	for _, opt := range options {
		opt(o)
	}

	requestBody := &request{
		OperatorID:       operatorID,
		Amount:           amount,
		UseLocalAmount:   useLocalAmount,
		CustomIdentifier: o.CustomIdentifier,
		RecipientEmail:   recipientEmail,
		SenderPhone:      &o.SenderPhone,
	}


	rb, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	payload := strings.NewReader(string(rb))
	method := "POST"
	client := c.HttpClient
	req, _ := http.NewRequest(method, c.BaseURL, payload)

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var e error
	var ar Transaction
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, err
		}
		return nil, e

	}

	err = json.NewDecoder(res.Body).Decode(&ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

