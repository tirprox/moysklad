package test

import (
	"github.com/tirprox/moysklad"
	"testing"
)

func TestGetOne(t *testing.T) {
	client := moysklad.NewClient("gleb@salonvbelom", "k#TfUG6v")
	params := moysklad.QueryParams{}
	params.AddQuery("limit", "100")

	storeId := "ca07a57b-9c9c-11e8-9ff4-34e800073881"

	item := client.GetStore(storeId, params)

	if item.ID != storeId {
		t.Errorf("Store ID is incorrect, got: %s, want: %s.", item.ID, storeId)
	}
}
