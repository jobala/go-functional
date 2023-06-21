package types

type Customer struct {
	Name       string
	Address    string
	Enterprise bool
}

type OrderItem struct {
	Name        string
	ItemNumber  int
	Quantity    int
	Price       int
	Backordered bool
}

type Order struct {
	Orderid         int
	ShippingAddress string
	Expedited       bool
	Customer        Customer
}

var Orders = []Order{Ord1, Ord2}

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
