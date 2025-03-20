package commands

import (
	"encoding/xml"

	"github.com/dracoDevs/go-ebay-plus/internal/ebay"
)

type AddItem struct {
	Currency              string
	Country               string
	DispatchTimeMax       int    `xml:",omitempty"`
	ConditionID           int    `xml:",omitempty"`
	Title                 string `xml:",omitempty"`
	Description           string `xml:",omitempty"`
	StartPrice            string
	BuyItNowPrice         string `xml:",omitempty"`
	ListingType           string `xml:",omitempty"`
	Quantity              uint   `xml:",omitempty"`
	PaymentMethods        string `xml:",omitempty"`
	PayPalEmailAddress    string `xml:",omitempty"`
	ListingDuration       string
	ShippingDetails       *ShippingDetails `xml:",omitempty"`
	PrimaryCategory       *PrimaryCategory
	Storefront            *Storefront            `xml:",omitempty"`
	PostalCode            string                 `xml:",omitempty"`
	ReturnPolicy          *ReturnPolicy          `xml:",omitempty"`
	PictureDetails        *PictureDetails        `xml:",omitempty"`
	ProductListingDetails *ProductListingDetails `xml:",omitempty"`
	ItemSpecifics         []ItemSpecifics        `xml:",omitempty"`
}

func (c AddItem) CallName() string {
	return "AddItem"
}

func (c AddItem) ParseResponse(r []byte) (ebay.EbayResponse, error) {
	var xmlResponse AddItemResponse
	err := xml.Unmarshal(r, &xmlResponse)

	return xmlResponse, err
}

func (c AddItem) Body() interface{} {
	type Item struct {
		AddItem
	}

	return Item{c}
}

type AddItemResponse struct {
	ebay.OtherEbayResponse

	ItemID string
}

func (r AddItemResponse) ResponseErrors() ebay.EbayErrors {
	return r.OtherEbayResponse.Errors
}
