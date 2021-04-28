# Account Operations

Use the account operations to perform account related actions

## Retrieve account balance info

```golang
package main

import (
	"fmt"
	"github.com/Reloadly/reloadly"
)

func main (){
	c, err := reloadly.NewClient("Nxcccccc9", "c7gxxxxxx", true)
	if err != nil {
		fmt.Println(err)
	}

	f, err := c.GetBalance()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f.Balance)
} 
```