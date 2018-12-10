package moysklad

import "github.com/tirprox/moysklad/codec"

func Store(item interface{}) codec.Store {
	return item.(codec.Store)
}
