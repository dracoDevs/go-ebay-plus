package ebay

import (
	"encoding/xml"
	"fmt"
	"time"
)

type ebayRequest struct {
	conf    EbayConf
	command Command
}


type EbayResponse interface {
	Failure() bool
	ResponseErrors() EbayErrors
}

type OtherEbayResponse struct {
	Timestamp time.Time
	Ack       string
	Errors    []ebayResponseError
}

type ebayResponseError struct {
	ShortMessage        string
	LongMessage         string
	ErrorCode           int
	SeverityCode        string
	ErrorClassification string
}

func (r OtherEbayResponse) Failure() bool {
	return r.Ack == "Failure"
}

func (r OtherEbayResponse) ResponseErrors() EbayErrors {
	return r.Errors
}

func (c ebayRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	startElement := xml.StartElement{
		Name: xml.Name{
			Space: "urn:ebay:apis:eBLBaseComponents",
			Local: fmt.Sprintf("%sRequest", c.command.CallName()),
		},
	}

	err := e.EncodeToken(startElement)

	if err != nil {
		return err
	}

	type RequesterCredentials struct {
		EbayAuthToken string `xml:"eBayAuthToken"`
	}

	err = e.Encode(
		RequesterCredentials{
			EbayAuthToken: c.conf.AuthToken,
		},
	)

	if err != nil {
		return err
	}

	err = e.Encode(c.command.Body())

	if err != nil {
		return err
	}

	endElement := xml.EndElement{
		Name: xml.Name{
			Space: "urn:ebay:apis:eBLBaseComponents",
			Local: fmt.Sprintf("%sRequest", c.command.CallName()),
		},
	}

	err = e.EncodeToken(endElement)

	if err != nil {
		return err
	}

	return nil
}