package main

import (
	"github.com/ploynomail/gonetworkmanager/v2"
)

func main() {
	c, err := gonetworkmanager.GetConnectionByName("br0")
	if err != nil {
		panic(err)
	}
	if c == nil {
		panic("Connection not found")
	}
}
