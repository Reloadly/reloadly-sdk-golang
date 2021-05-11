# Topup Operations

In order to send a successful topup. There are a few prerequisites to the system. We need to know the phone number to
send the topup to, the operator id of the phone number, the country of the operator, the amount for the topup.

## Topups - Send topup to phone

```golang
package main

import (
	"fmt"
	reloadly "github.com/reloadly/reloadly-sdk-golang/airtime"
)

func main (){
	c, err := reloadly.NewClient("NxhvFxDrXtX34VL1DLHsbQGLOPgE8qL9", "c7gcNR4dOIAlmalojV35okbq--tqqVhPPr2bvcsXlmKadewvaXdRSxENwfpyFywh", true)
	if err != nil {
		fmt.Println(err)
	}

	rPhone := reloadly.Phone{
		Number:      "+2340123456",
		CountryCode: "NG",
	}

	t, err := c.Topup("100","341", false, rPhone)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Topup",t)
}
```

## Topup Options - 

You can futher customize the bahaviour of this library as it interacxts with the reloadly API using functional options provided. If for example, You intend to add a custom identifier for the topup transaction,you can customize it as seen below
```golang
    customIdentifier := reloadly.AddCustomIdentifier("First Transaction")
    _, err := c.Topup("100","341", false, rPhone, customIdentifier)
    if err != nil {
        fmt.Println(err)
    }
```
This library supports all optional parameters specified in the [official docs](https://developers.reloadly.com/api.html#send-topup). The name of the optional functions usually start with a verb such as ```addCustomIdentifier``` 
## Topups - Send Nauta Cuba topups

Reloadly also supports Nauta Cuba for top-ups. However the process is a bit different from sending phone topups. Instead
of using the ```Phone``` struct, use the recipient email and that's it. The rest of the process is exactly the same as sending any other topup.

Note, There are two types of email domains that are allowed for Nauta Cuba Top-up : ```@nauta.com.cu```
and ```@nauta.co.cu```

```golang
package main

import (
	"fmt"
	"github.com/Ghvstcode/reloadly"
)

func main (){
	c, err := reloadly.NewClient("NxhvFxDrXtX34VL1DLHsbQGLOPgE8qL9", "c7gcNR4dOIAlmalojV35okbq--tqqVhPPr2bvcsXlmKadewvaXdRSxENwfpyFywh", true)
	if err != nil {
		fmt.Println(err)
	}

	t, err := c.NautaCubaTopup("100","341", false, "test@nauta.co.cu")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Topup",t)
}
```