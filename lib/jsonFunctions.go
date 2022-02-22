package lib

import "encoding/json"

func LoadProductsJson(d1 []byte) Products {
	var products Products
	json.Unmarshal(d1, &products)
	return products
}

func LoadOrdersJson(d1 []byte) Orders {
	var orders Orders
	json.Unmarshal(d1, &orders)
	return orders
}


func ToJson(d1 interface{}) string {
	jsondata, _ := json.Marshal(d1)
	return string(jsondata)
}