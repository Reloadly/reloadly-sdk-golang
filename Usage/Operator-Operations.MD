# Operators Operations

Apart from supporting Over 140 Countries, Reloadly also supports 600+ Operators. The SDK operators options allows for
the retrieval complete detail of each operator, including what type of operator this is, what topup types it support and
even details on the commissions for the operator.

Within the reloadly platform, There exists two types of Operators. One that support Range values (Anything between the
minimum and maximum range). While the other that support Fixed values (Only a certain values are supported). Reloadly
will return you the type of the operator within the response in denominationType variable. If this is set to ```RANGE```
you will receive the minimum and maximum values in the minAmount and maxAmount variables for that operator. However, if
the denomination type is ```FIXED``` you will not get these values but rather get an array of all values supported in
the fixedAmounts variable. **Now a point to remember here is that these values are already converted into your account's
currency**.

## Operators - List

```golang
    package main
    
    import (
        "fmt"
        "github.com/Ghvstcode/reloadly"
    )
    
    func main (){
        c, err := reloadly.NewClient("Nxhxxxxxxx", "c7xxxxx", true)
        if err != nil {
            fmt.Println(err)
        }
    
        f, err := c.GetOperators()
		if err != nil {
			fmt.Println(err)
		}
        fmt.Println(f)
    }       
```

## Operators - List with filters

Additional operator filters may be enabled

```golang
    ------
    AddPageSize := reloadly.AddPageSize(2)
    AddPin := reloadly.AddPin(true)
    f, err := c.GetOperators(AddPageSize, AddPin)
    if err != nil {
        fmt.Println(err)
    }
```

By default all the above filters are set to ```true``` except for ```includeSuggestedAmounts```
and ```includeSuggestedAmountsMap```

## Operators - Get by id

```golang
    func main (){
        c, err := reloadly.NewClient("Nxhxxxxxxx", "c7xxxxx", true)
        if err != nil {
        fmt.Println(err)
        }
        
        f, err := c.GetOperatorsById(1)
        if err != nil {
        fmt.Println(err)
        }
        fmt.Println(f)
    }
```

## Operators - Get by country

```golang
    func main (){
    	c, err := reloadly.NewClient("Nxhxxxxxxx", "c7xxxxx", true)
        if err != nil {
            fmt.Println(err)
        }
        
        f, err := c.GetOperatorsByISO("NG")
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(f)
    }
```

## Operators - Auto-detect

Reloadly also provide a simple way to automatically detect the operator for any given number.

```golang
    f, err := c.GetOperatorsByPhone("+234xxxxxx", "NG")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(f)
```

## Operators - Auto-detect with filters

```golang
    --------
    suggestedAmount := reloadly.AddSuggestedAmounts(true)
    suggestedAmountMap := reloadly.AddSuggestedAmountsMap(false)
    f, err := c.GetOperatorsByPhone("+234XXXXXX", "NG", suggestedAmount, suggestedAmountMap)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(f)
```

## Operators - Calculate FX rate

In order to estimate what amount will be received on the receiver end. For example, If your account is in US Dollar and
you are trying to send a transaction to a nigerian operator, you can quickly calculate what amount you will receive in
Nigerian Naira.

```golang
    --------
    f, err := c.GetFxRate(1,1)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(f)
```