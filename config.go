package main

import "github.com/ekomobile/dadata/v2/client"

type RequestAddressSearch struct {
	Query string `json:"query"`
}

type ResponseAddress struct {
	Addresses []*Address `json:"addresses"`
}

type RequestAddressGeocode struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}
type Address struct {
	Country      string `json:"country"`
	Region       string `json:"region"`
	Area         string `json:"area"`
	City         string `json:"city"`
	CityDistrict string `json:"city_district"`
	Settlement   string `json:"settlement"`
	Street       string `json:"street"`
	House        string `json:"house"`
	Block        string `json:"block"`
	Flat         string `json:"flat"`
	PostalBox    string `json:"postal_box"`
	Timezone     string `json:"timezone"`
}

var keys = client.Credentials{
	ApiKeyValue:    "007233085e9e9f5e7fa5871f0828d87ff737f0bf",
	SecretKeyValue: "e4f919962e435bc587ac8c21f5e1b210c7bf9b11",
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
