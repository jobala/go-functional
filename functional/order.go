package functional

import "github.com/jobala/functional-go/types"

func getCustomerName(order types.Order) string {
	return order.Customer.Name
}

func getCustomerAddress(order types.Order) string {
	return order.Customer.Address
}

func getShippingAddress(order types.Order) string {
	return order.ShippingAddress
}

func isExpedited(order types.Order) bool {
	return order.Expedited
}

type predicateFunc func(types.Order) bool
type attrFunc func(types.Order) string

func getFilteredInfo(predicate predicateFunc, f attrFunc, orders []types.Order) []string {
	var res []string

	for _, order := range orders {
		if predicate(order) {
			res = append(res, f(order))
		}
	}

	return res
}

func GetExpeditedOrderCustomerNames(orders []types.Order) []string {
	return getFilteredInfo(isExpedited, getCustomerName, orders)
}

func GetExpeditedOrdersCustomerAddresses(orders []types.Order) []string {
	return getFilteredInfo(isExpedited, getCustomerAddress, orders)
}

func GetExpeditedOrdersShippingAddresses(orders []types.Order) []string {
	return getFilteredInfo(isExpedited, getShippingAddress, orders)
}
