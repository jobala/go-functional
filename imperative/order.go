package imperative

var Orders = []Order{Ord1, Ord2}

func (o Order) GetExpeditedOrderCustomerNames() []string {
	res := make([]string, 0)

	for _, order := range Orders {
		if order.Expedited {
			res = append(res, order.Customer.Name)
		}
	}

	return res
}

func (o Order) GetExpeditedOrdersCustomerAddresses() []string {
	res := make([]string, 0)

	for _, order := range Orders {
		if order.Expedited {
			res = append(res, order.Customer.Address)
		}
	}

	return res
}

func (o Order) GetExpeditedOrdersShippingAddresses() []string {
	res := make([]string, 0)

	for _, order := range Orders {
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

var cus1 = Customer{
	Name:       "Heart of Gold",
	Address:    "The Milky Way Galaxy",
	Enterprise: false,
}

var cus2 = Customer{
	Name:       "Milliways Restaurant",
	Address:    "Magrathea",
	Enterprise: true,
}

var Ord1 = Order{
	Customer:        cus1,
	Expedited:       false,
	ShippingAddress: "Infinitely Improbable",
}

var Ord2 = Order{
	Customer:        cus2,
	Expedited:       true,
	ShippingAddress: "Magrathea",
}
