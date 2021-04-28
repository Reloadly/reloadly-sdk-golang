# Countries Operations

Reloadly supports 140+ Countries around the globe. You can get a list of all or specific supported countries. The
response will give you a list with complete details, iso, flag as well as calling codes for each country. You can also
further filter the countries by getting details for a specific country by its ISO-Alpha2 code.
See https://www.nationsonline.org/oneworld/country_code_list.htm for more details regarding country ISO codes.

## Countries - List

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

	f, err := c.GetCountries()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
} 
```

## Countries - Get by country ISO code

```golang
    --------
    f, err := c.GetCountriesByIso("NG")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
```
