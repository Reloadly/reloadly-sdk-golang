package airtime

import (
	"encoding/base64"
	"encoding/json"
	"github.com/reloadly/reloadly-sdk-golang/authentication"
	error2 "github.com/reloadly/reloadly-sdk-golang/error"
	"net/http"
	"strconv"
)

var (
	baseUrl string
)

//Client is a struct that represents the necessary Data for interacting with the Reloadly API through this SDk.
type Client struct {
	ClientID string
	ClientSecret string
	SandBox bool
	HttpClient HTTPClient
	AuthHeader string
	BaseURL string
	telemetryHeader string
}


type telemetry struct {
	APIVersion string `json:"api-version"`
	Name       string `json:"name"`
	Env        struct {
		Golang            string `json:"golang"`
		ReloadlySdkGolang string `json:"reloadly-sdk-golang"`
	} `json:"env"`
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

//ClientOptions is used to further customize the Client that would be used in interacting with this SDK
type ClientOptions struct {
	//Telemetry is used to specify if Telemetry should be enabled or disabled. By default, It is enabled.
	Telemetry bool
	//HTTPClient Can be provided to the Reloadly SDK.
	HTTPClient HTTPClient
}

//ClientOpts are options used to customise the Client.
type ClientOpts func(opts *ClientOptions)

//NewClient creates a new Client for interacting with this SDK. Under the hood, it uses the authentication module to authenticate users.
func NewClient(clientId, clientSecret string, sandbox bool, opts ...ClientOpts)(*Client, error){
	var telemetryHeader string
	var HTTPClient HTTPClient

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

	//Configure telemetry
	t := &telemetry{
		APIVersion: "application/com.reloadly.topups-v1+json",
		Name:       "reloadly-sdk-golang",
		Env: struct {
			Golang            string `json:"golang"`
			ReloadlySdkGolang string `json:"reloadly-sdk-golang"`
		}{
			Golang: "1.15.5",
			ReloadlySdkGolang: "1.0.0",
		},
	}
	telemetryJson, _ := json.Marshal(t)


	//Check & Init configurable options
	o := &ClientOptions{}
	for _, opt := range opts {
		opt(o)
	}


	telemetryHeader = ""
	if o.Telemetry {
		telemetryHeader = base64.StdEncoding.EncodeToString(telemetryJson)
	}

	HTTPClient = &http.Client{}
	if o.HTTPClient != nil{
		HTTPClient = o.HTTPClient
	}

	authHeader := accessResponse.AccessToken
	//Return Created Client
	return &Client{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		SandBox:      sandbox,
		HttpClient: HTTPClient,
		AuthHeader: "Bearer " + authHeader,
		BaseURL: baseUrl,
		telemetryHeader: telemetryHeader,
	}, nil
}

//ConfigureHTTP is used to configure the HTTP client used to make request to the Reloadly API.
func (c *Client) ConfigureHTTP(h *http.Client) *Client{
	c.HttpClient = h
	return c
}


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

