package imperative

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
