package main

import (
	"fmt"
	"github.com/Ghvstcode/reloadly"
)

func main (){
	c, err := reloadly.NewClient("Nxhvxxxx", "xxxxxxh", true)
	if err != nil {
		fmt.Println(err)
	}
	AddPageSize := reloadly.AddPageSize(2)
	AddPin := reloadly.AddPin(true)
	f, err := c.GetOperators(AddPageSize, AddPin)
	fmt.Println(f)
	fmt.Println(err)
}
