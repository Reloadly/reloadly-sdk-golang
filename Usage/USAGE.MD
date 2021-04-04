# Airtime API

The implementation is based on the [Airtime API Docs](https://developers.reloadly.com/#airtime-api).

## Usage

Create an `AirtimeAPI` instance by providing the Application details (client id & secret) from
the [dashboard](https://www.reloadly.com/developers/api-settings).

Some key things to keep in mind regarding the Airtime API :

* The API has 2 environments, SANDBOX (for development & testing) and LIVE.
* Each environment has a set of credentials (client id & secret) that are different from the other.<br />
    * SANBOX credentials can only be used for SANDBOX environment
    * LIVE credentials can only be used for LIVE environment
* You MUST supply either the credentials, or an access token in order to call the API
  <br /><br />

As stated above, requests to the Airtime API require authentication/authorization, there are several options :

### Option 1

Set the client id, client secret & environment; The most straight-forward or simplest way is by calling the NewClient function. The NewClient function returns a Client struct which contains the authentication header(formed from the accesstoken) and other necessary components for interacting with the API through this SDK.

```golang
    package main

    import (
      	"github.com/Ghvstcode/reloadly" 
      	"fmt"
    )

    func main(){
      client, err := reloadly.NewClient(xx1234, xx23, false)
      if err != nil {
        fmt.Println(err)
      }
    }

    
```

### Option 2

You may alternatively acquire an access token from the
[AuthenticationAPI](https://github.com/reloadly/reloadly-sdk-java/blob/master/reloadly-java-sdk-authentication/USAGE.md)
and then set it.

```golang
      package main

      import (
        "github.com/Ghvstcode/reloadly/authentication"
        "fmt"
      )

      func main(){
        authClient, err := authentication.NewAuthClient(xx1234, xx23, false)
        if err != nil {
          fmt.Println(err)
        }
        
        token, err := authClient.GetAccessToken()
        if err != nil {
          fmt.Println(err)
        }
        
        fmt.Println(token.AccessToken)
      }
```

**Note : Access tokens obtain for Reloadly APIs have a finite
lifetime. [See the API docs](https://developers.reloadly.com/#authentication_auth_anc)**

### Customizing The API Client Instance
You can customize the Client Instance used in communicating with the Reloadly API to fit your needs
```golang
    httpClient := http.Client{
        Timeout: 5 * time.second
    }

    airtimeClient := client.ConfigureHTTP(httpClient)
```