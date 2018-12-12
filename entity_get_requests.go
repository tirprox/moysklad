package moysklad

import (
	"github.com/mitchellh/mapstructure"
	"github.com/tirprox/moysklad/codec"
)

func (c moyskladClient) GetStore(id string, params QueryParams) (entity codec.Store) {

	item := c.GetOne("store", id, params)
	mapstructure.Decode(item, &entity)

	return
}

func (c moyskladClient) GetStores(params QueryParams) (entities []codec.Store) {
	items := c.GetMany("store", params)

	for _, item := range items {
		entity := codec.Store{}
		mapstructure.Decode(item, &entity)
		entities = append(entities, entity)

	}
	return
}

func (c moyskladClient) GetAssortment(params QueryParams) (products []codec.Product, variants []codec.Variant) {
	items := c.GetMany("assortment", params)

	for _, item := range items {
		meta := extractMeta(item)
		entityType := meta["type"]

		switch entityType {
		case "product":
			product := codec.Product{}
			mapstructure.Decode(item, &product)
			products = append(products, product)
		case "variant":
			variant := codec.Variant{}
			mapstructure.Decode(item, &variant)
			variants = append(variants, variant)
		}
		//fmt.Println(entityType)
	}

	return
}

func (c moyskladClient) GetCounterparty(id string, params QueryParams) (entity codec.Counterparty) {
	item := c.GetOne("counterparty", id, params)
	mapstructure.Decode(item, &entity)
	return
}

func (c moyskladClient) GetCounterparties(params QueryParams) (entities []codec.Counterparty) {
	items := c.GetMany("counterparty", params)

	for _, item := range items {
		entity := codec.Counterparty{}
		mapstructure.Decode(item, &entity)
		entities = append(entities, entity)

	}
	return
}

func (c moyskladClient) GetProduct(id string, params QueryParams) (entity codec.Product) {
	item := c.GetOne("product", id, params)
	mapstructure.Decode(item, &entity)
	return
}

func (c moyskladClient) GetProducts(params QueryParams) (entities []codec.Product) {
	items := c.GetMany("product", params)

	for _, item := range items {
		entity := codec.Product{}
		mapstructure.Decode(item, &entity)
		entities = append(entities, entity)

	}
	return
}

func (c moyskladClient) GetVariant(id string, params QueryParams) (entity codec.Product) {
	item := c.GetOne("variant", id, params)
	mapstructure.Decode(item, &entity)
	return
}

func (c moyskladClient) GetVariants(params QueryParams) (entities []codec.Variant) {
	items := c.GetMany("variant", params)

	for _, item := range items {
		entity := codec.Variant{}
		mapstructure.Decode(item, &entity)
		entities = append(entities, entity)

	}
	return
}

func (c moyskladClient) GetCustomerOrder(id string, params QueryParams) (entity codec.Customerorder) {

	item := c.GetOne("customerorder", id, params)
	mapstructure.Decode(item, &entity)
	return
}

func (c moyskladClient) GetCustomerOrders(params QueryParams) (entities []codec.Customerorder) {
	items := c.GetMany("customerorder", params)

	for _, item := range items {
		entity := codec.Customerorder{}
		mapstructure.Decode(item, &entity)
		entities = append(entities, entity)

	}
	return
}
