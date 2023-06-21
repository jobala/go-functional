package main

import (
	"fmt"

	"github.com/jobala/functional-go/functional"
	"github.com/jobala/functional-go/types"
)

func main() {
	for _, name := range functional.GetExpeditedOrderCustomerNames(types.Orders) {
		fmt.Println(name)
	}

	for _, address := range functional.GetExpeditedOrdersCustomerAddresses(types.Orders) {
		fmt.Println(address)
	}

	for _, address := range functional.GetExpeditedOrdersShippingAddresses(types.Orders) {
		fmt.Println(address)
	}
}
