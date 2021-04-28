# Report Operations

Retrieve various reports such as transaction history etc...

## Reports - List transaction history

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
    
        t, err := c.GetTransactions()
        if err != nil {
            fmt.Println(err)
        }
    
        fmt.Println("Transactions",t)
    }
```

## Reports - List transaction history with filters

```golang
---------------
    page := reloadly.FilterByPage(2)
    size := reloadly.FilterBySize(4)
    f, err := c.GetTransactions(page, size)
    if err != nil {
        fmt.Println(err)
    }
```

## Reports - Get transaction history by id

```golang
    --------
    f, err := c.GetTransactionsByID(74)
    if err != nil {
     fmt.Println(err)
    }
```