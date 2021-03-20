package reloadly

import (
	"github.com/Ghvstcode/reloadly/authentication"
	"net/http"
	"strconv"
)

var baseUrl string

type Client struct {
	ClientID string
	ClientSecret string
	SandBox bool
	HttpClient http.Client
	AuthHeader string
	BaseURL string
}

type FilterOptions struct{
	Page string
	Size string
}

type Filters func(opts *FilterOptions)


func NewClient(clientId, clientSecret string, sandbox bool)(*Client, error){
	if clientId == ""{
		return nil, &Error{Message: "INVALID_CREDENTIALS"}
	}

	authClient, err := authentication.NewAuthClient(clientId, clientSecret, sandbox)
	if err != nil {
		return nil, &Error{Message: err.Error()}
	}

	if sandbox{
		baseUrl = "https://topups-sandbox.reloadly.com"
	} else {
		baseUrl = "https://topups.reloadly.com"
	}

	//Retrieve accessToken
	accessResponse, err := authClient.GetAccessToken()
	if err != nil {
		return nil, &Error{Message: err.Error()}
	}

	authHeader := accessResponse.AccessToken

	//Return Created Client
	return &Client{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		SandBox:      sandbox,
		HttpClient: *http.DefaultClient,
		AuthHeader: "Bearer " + authHeader,
		BaseURL: baseUrl,
	}, nil
}

//func ()RefreshClient(){
//
//}

func FilterByPage(page int) Filters {
	return func(s *FilterOptions) {
		s.Page = strconv.Itoa(page)
	}
}

func FilterBySize(page int) Filters {
	return func(s *FilterOptions) {
		s.Page = strconv.Itoa(page)
	}
}

