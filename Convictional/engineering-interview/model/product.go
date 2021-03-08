package model

import "encoding/json"

type Product struct {
	ID        json.Number
	Title     string
	Vendor    string
	Body_html string
	Variants  []Variant
	Images    []Image
}

type Variant struct {
	id                 string
	title              string
	sku                string
	available          bool
	inventory_quantity int
}

type Image struct {
	source    string
	varientId string
}
