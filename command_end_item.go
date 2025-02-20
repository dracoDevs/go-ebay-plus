package ebay

import (
	"encoding/xml"
)

type EndItem struct {
	ItemID       string       `xml:"ItemID"`
	EndingReason EndingReason `xml:"EndingReason"`
}

type EndingReason string

const (
	CustomCode        EndingReason = "CustomCode"
	Incorrect         EndingReason = "Incorrect"
	LostOrBroken      EndingReason = "LostOrBroken"
	NotAvailable      EndingReason = "NotAvailable"
	OtherListingError EndingReason = "OtherListingError"
	ProductDeleted    EndingReason = "ProductDeleted"
	SellToHighBidder  EndingReason = "SellToHighBidder"
	Sold              EndingReason = "Sold"
)

func (c EndItem) CallName() string {
	return "EndItem"
}

func (c EndItem) Body() interface{} {
	return EndItem{
		ItemID: c.ItemID,
		EndingReason: c.EndingReason,
	}
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
