package lib

import "encoding/json"

func LoadJson(d1 []byte) Orders {
	var orders Orders
	json.Unmarshal(d1, &orders)
	return orders
}


func ToJson(d1 interface{}) string {
	jsondata, _ := json.Marshal(d1)
	return string(jsondata)
}