package authentication

import (
	"bytes"
	"encoding/json"
	"errors"
	error2 "github.com/reloadly/reloadly-sdk-golang/error"
	"net/http"
)

type AuthClient struct {
	ClientID string
	ClientSecret string
	SandBox bool
	BearerToken string
	HttpClient *http.Client
	URL string
}

//AccessResponse is the Struct that serializes the JSON response gotten when a Request for an AccessToken is made!
//Along with the AccessToken,it contains more info about the AccessToken, like the Scope, TokenType & ExpiresIn.
type AccessResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

//NewAuthClient Creates an AuthenticationAPI instance by providing the Application credentials details (client id & secret) from the dashboard.
//The created Client is used to interact with this authentication sub-module.
func NewAuthClient(clientId, clientSecret string, sandbox bool)(*AuthClient, error){
	if clientId == "" || clientSecret == ""{
		return nil, errors.New("INVALID_CREDENTIALS")
	}
	return &AuthClient{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		SandBox:      sandbox,
		HttpClient: http.DefaultClient,
		URL: "https://auth.reloadly.com/oauth/token",
	}, nil
}

//GetAccessToken returns an AccessResponse struct which contains the AccessToken and related information which can be used for making authenticated requests
func (a *AuthClient) GetAccessToken()(*AccessResponse,error){
	method := "POST"

	environment := "https://topups-sandbox.reloadly.com"
	if !a.SandBox{
		environment = "https://topups.reloadly.com"
	}

	p := map[string]string{
		"client_id": a.ClientID,
		"client_secret": a.ClientSecret,
		"grant_type": "client_credentials",
		"audience": environment,
	}

	data, _ := json.Marshal(p)

	client := a.HttpClient
	req, _ := http.NewRequest(method, a.URL, bytes.NewBuffer(data))
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var e error2.ErrorResponse
	var ar AccessResponse
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

//ConfigureHTTP is Used to configure additional options, connect and read timeouts for the HTTP client.
func (a *AuthClient) ConfigureHTTP(h *http.Client){
	a.HttpClient = h
}