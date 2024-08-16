package entity

type Address struct {
	Street  string     `json:"street"`
	Suite   string     `json:"suite"`
	City    string     `json:"city"`
	Zipcode string     `json:"zipcode"`
	Geo     AddressGeo `json:"geo"`
}
