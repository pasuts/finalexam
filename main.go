package main

import (
	"github.com/pasuts/goonline/customer"
)

func main() {
	r := customer.SetupRouter()
	r.Run(":2019")
}
