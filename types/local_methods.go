package types

import "encoding/json"

func (base * RSS3Base) ToJson() []byte {
	baseJson, _ := json.Marshal(&base)
	return baseJson
}

func (persona * RSS3) ToJson() []byte {
	personaJson, _ := json.Marshal(&persona)
	return personaJson
}

func (items * RSS3Items) ToJson() []byte {
	itemsJson, _ := json.Marshal(&items)
	return itemsJson
}

func (list * RSS3List) ToJson() []byte {
	listJson, _ := json.Marshal(&list)
	return listJson
}

func (item * RSS3Item) ToJson() []byte {
	itemJson, _ := json.Marshal(&item)
	return itemJson
}
