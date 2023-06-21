package main

import (
	"fmt"

	"github.com/jobala/functional-go/imperative"
)

func main() {
	for _, name := range imperative.Ord1.GetExpeditedOrderCustomerNames() {
		fmt.Println(name)
	}

	for _, address := range imperative.Ord1.GetExpeditedOrdersCustomerAddresses() {
		fmt.Println(address)
	}

	for _, address := range imperative.Ord1.GetExpeditedOrdersShippingAddresses() {
		fmt.Println(address)
	}
}
