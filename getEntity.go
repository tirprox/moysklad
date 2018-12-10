package moysklad

import "github.com/tirprox/moysklad/codec"

/* Signature looks like this:
Single entity:
	id string
	filter map[string]string
	query map[string]string

All entities:
	filter map[string]string
	query map[string]string

*/

func GetStore() (store codec.Store) {

	return
}

func GetStores() (stores []codec.Store) {

	return
}

func GetAssortment() (products []codec.Product, variants []codec.Variant) {

	return
}

func GetCounterparty() (counterparty codec.Counterparty) {

	return
}

func GetCounterparties() (counterparties []codec.Counterparty) {

	return
}

func GetProduct() (product codec.Product) {

	return
}

func GetProducts() (products []codec.Product) {

	return
}

func GetCustomerOrder() (customerorder codec.Customerorder) {

	return
}

func GetCustomerOrders() (customerorders []codec.Customerorder) {

	return
}
