You can use the authentication module to alternatively acquire authentication details(access token) from the
[AuthenticationAPI](https://developers.reloadly.com/#authentication_auth_anc)
and then set it.

```golang
      package main

      import (
        "github.com/reloadly/reloadly-sdk-golang/authentication"
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