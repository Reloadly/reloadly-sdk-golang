package reloadly

import (
	"github.com/Ghvstcode/reloadly/authentication"
	error2 "github.com/Ghvstcode/reloadly/error"
	"net/http"
	"strconv"
)

var baseUrl string

//Client is a struct that represents the necessary Data for interacting with the Reloadly API through this SDk.
type Client struct {
	ClientID string
	ClientSecret string
	SandBox bool
	HttpClient HTTPClient
	AuthHeader string
	BaseURL string
}

//HTTPClient is an interface with the Do function for interacting with the Relaodly API! You do not have to interact with this interface, It is implemented by default.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

//FilterOptions is a struct representing pagination filters used across this library.
type FilterOptions struct{
	Page string
	Size string
}

//Filters is a type representing pagination filters used across this project.
type Filters func(opts *FilterOptions)


func NewClient(clientId, clientSecret string, sandbox bool)(*Client, error){
	if clientId == ""{
		return nil, &error2.ErrorResponse{Message: "INVALID_CREDENTIALS"}
	}

	authClient, err := authentication.NewAuthClient(clientId, clientSecret, sandbox)
	if err != nil {
		return nil, &error2.ErrorResponse{Message: err.Error()}
	}

	if sandbox{
		baseUrl = "https://topups-sandbox.reloadly.com"
	} else {
		baseUrl = "https://topups.reloadly.com"
	}

	//Retrieve accessToken
	accessResponse, err := authClient.GetAccessToken()
	if err != nil {
		return nil, &error2.ErrorResponse{Message: err.Error()}
	}

	authHeader := accessResponse.AccessToken

	//Return Created Client
	return &Client{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		SandBox:      sandbox,
		HttpClient: &http.Client{},
		AuthHeader: "Bearer " + authHeader,
		BaseURL: baseUrl,
	}, nil
}

//ConfigureHTTP is used to configure the HTTP client used to make request to the Reloadly API.
func (c *Client) ConfigureHTTP(h *http.Client) *Client{
	c.HttpClient = h
	return c
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

