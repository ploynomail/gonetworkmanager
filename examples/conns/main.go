package main

import (
	"fmt"

	"github.com/ploynomail/gonetworkmanager/v2"
)

func main() {
	c, err := gonetworkmanager.GetConnectionByName("bond0")
	if err != nil {
		panic(err)
	}
	if c == nil {
		panic("Connection not found")
	}
	st, _ := c.GetSettings()
	fmt.Println(st)
}
