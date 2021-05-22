package methods

import (
	"encoding/json"
	"github.com/nyawork/rss3go_lib/types"
)

func Json2RSS3Base(baseJson []byte) types.RSS3Base {
	var base types.RSS3Base
	_ = json.Unmarshal(baseJson, &base)

	return base
}

func Json2RSS3(baseJson []byte) types.RSS3 {
	var persona types.RSS3
	_ = json.Unmarshal(baseJson, &persona)

	return persona
}

func Json2RSS3Items(itemsJson []byte) types.RSS3Items {
	var items types.RSS3Items
	_ = json.Unmarshal(itemsJson, &items)

	return items
}

func Json2RSS3List(linkJson []byte) types.RSS3List {
	var list types.RSS3List
	_ = json.Unmarshal(linkJson, &list)

	return list
}

func Json2RSS3Item(itemsJson []byte) types.RSS3Item {
	var item types.RSS3Item
	_ = json.Unmarshal(itemsJson, &item)

	return item
}

