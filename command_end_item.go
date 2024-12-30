package ebay

import (
	"encoding/xml"
	"errors"
)

type EndItem struct {
	ItemID       string
	EndingReason string
}

const (
	CustomCode        string = "CustomCode"
	Incorrect         string = "Incorrect"
	LostOrBroken      string = "LostOrBroken"
	NotAvailable      string = "NotAvailable"
	OtherListingError string = "OtherListingError"
	ProductDeleted    string = "ProductDeleted"
	SellToHighBidder  string = "SellToHighBidder"
	Sold              string = "Sold"
)

var validReasons = map[string]bool{
	CustomCode:        true,
	Incorrect:         true,
	LostOrBroken:      true,
	NotAvailable:      true,
	OtherListingError: true,
	ProductDeleted:    true,
	SellToHighBidder:  true,
	Sold:              true,
}

func (c EndItem) CallName() string {
	return "EndItem"
}

func (c EndItem) Body() (interface{}, error) {
	if !validReasons[c.EndingReason] {
		return nil, errors.New("invalid ending reason")
	}
	type Item struct {
		EndItem
	}
	return Item{c}, nil
}

func (c EndItem) ParseResponse(r []byte) (EbayResponse, error) {
	var xmlResponse EndItemResponse
	err := xml.Unmarshal(r, &xmlResponse)
	return xmlResponse, err
}

type EndItemResponse struct {
	ebayResponse

	EndTime string `xml:"EndTime"`
}

func (r EndItemResponse) ResponseErrors() ebayErrors {
	return r.ebayResponse.Errors
}
