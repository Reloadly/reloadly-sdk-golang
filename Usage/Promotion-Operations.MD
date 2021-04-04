# Promotions Operations

Reloady also support operators promotions. These are provided by the operators and can be activated by sending a
specific topup amount as per the details of the promotion. Using the promotion operations you can retrieve all details
on the different operators promotions and to showcase these to your customers.

## Promotions - List

```golang
    package main
    
    import (
        "fmt"
        "github.com/Ghvstcode/reloadly"
    )
    
    func main (){
        c, err := reloadly.NewClient("XXXXX12", "XXXXX1233", true)
        if err != nil {
            fmt.Println(err)
        }
        f, err := c.GetPromotions()
		if err != nil {
			fmt.Println(err)
		}
        fmt.Println(f)
    }
```

## Promotions - List with filters

```golang
    ----------
    page := reloadly.FilterByPage(2)
    size := reloadly.FilterBySize(4)
    _, err := c.GetPromotions(page, size)
    if err != nil {
        fmt.Println(err)
    }

```

## Promotions - Get by id

```golang
    p, err := c.GetPromotionsById(180)
    if err != nil {
        fmt.Println(err)
    }
```

## Promotions - Get by operator id

```golang
    p, err := c.GetPromotionsByOperatorId(1554)
    if err != nil {
        fmt.Println(err)
    }
```

## Promotions - Get by country code

```golang
    p, err := c.GetPromotionsByCountryCode("US")
    if err != nil {
        fmt.Println(err)
    }
```