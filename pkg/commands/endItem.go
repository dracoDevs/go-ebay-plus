package commands

import (
	"encoding/xml"

	"github.com/dracoDevs/go-ebay-plus/internal/ebay"
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

func (c EndItem) ParseResponse(r []byte) (ebay.EbayResponse, error) {
	var xmlResponse EndItemResponse
	err := xml.Unmarshal(r, &xmlResponse)
	return xmlResponse, err
}

type EndItemResponse struct {
	ebay.OtherEbayResponse

	EndTime string `xml:"EndTime"`
}

func (r EndItemResponse) ResponseErrors() ebay.EbayErrors {
	return r.OtherEbayResponse.Errors
}
