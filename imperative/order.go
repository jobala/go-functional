package imperative

import "github.com/jobala/functional-go/types"

func (o Order) GetExpeditedOrderCustomerNames() []string {
	res := make([]string, 0)

	for _, order := range types.Orders {
		if order.Expedited {
			res = append(res, order.Customer.Name)
		}
	}

	return res
}

func (o Order) GetExpeditedOrdersCustomerAddresses() []string {
	res := make([]string, 0)

	for _, order := range types.Orders {
		if order.Expedited {
			res = append(res, order.Customer.Address)
		}
	}

	return res
}

func (o Order) GetExpeditedOrdersShippingAddresses() []string {
	res := make([]string, 0)

	for _, order := range types.Orders {
		if order.Expedited {
			res = append(res, order.ShippingAddress)
		}
	}

	return res
}

type Order struct {
	Orderid         int
	ShippingAddress string
	Expedited       bool
	Customer        Customer
}
