package ebay

import (
	"encoding/xml"
	"fmt"
)

type ebayRequest struct {
	conf    EbayConf
	command Command
}

func (c ebayRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	startElement := xml.StartElement{
		Name: xml.Name{
			Space: "urn:ebay:apis:eBLBaseComponents",
			Local: fmt.Sprintf("%sRequest", c.command.CallName()),
		},
	}

	if err := e.EncodeToken(startElement); err != nil {
		return err
	}

	type RequesterCredentials struct {
		EbayAuthToken string `xml:"eBayAuthToken"`
	}

	if err := e.Encode(RequesterCredentials{EbayAuthToken: c.conf.AuthToken}); err != nil {
		return err
	}

	if c.command.CallName() == "EndItem" {
		endItem := c.command.Body().(EndItem)
		if err := e.Encode(EndItem{
			ItemID:       endItem.ItemID,
			EndingReason: endItem.EndingReason,
		}); err != nil {
			return err
		}
	} else {
		if err := e.Encode(c.command.Body()); err != nil {
			return err
		}
	}

	endElement := xml.EndElement{
		Name: xml.Name{
			Space: "urn:ebay:apis:eBLBaseComponents",
			Local: fmt.Sprintf("%sRequest", c.command.CallName()),
		},
	}

	if err := e.EncodeToken(endElement); err != nil {
		return err
	}

	return nil
}
