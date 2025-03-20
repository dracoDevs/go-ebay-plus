package ebay

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/dracoDevs/go-ebay-plus/internal/utils"
)

type Command interface {
	Body() interface{}
	CallName() string
	ParseResponse([]byte) (EbayResponse, error)
}

type EbayConf struct {
	baseUrl string

	DevId, AppId, CertId string
	RuName, AuthToken    string
	SiteId               int
	Logger               func(...interface{})
}

func (e EbayConf) Sandbox() EbayConf {
	e.baseUrl = "https://api.sandbox.ebay.com"
	return e
}

func (e EbayConf) Production() EbayConf {
	e.baseUrl = "https://api.ebay.com"
	return e
}

func (e EbayConf) RunCommand(c Command) (EbayResponse, error) {
	ec := ebayRequest{conf: e, command: c}

	body := new(bytes.Buffer)
	body.Write([]byte(xml.Header))

	if err := xml.NewEncoder(body).Encode(ec); err != nil {
		return OtherEbayResponse{}, err
	}

	if c.CallName() == "EndItem" {
		bodyStr := utils.RemoveTagXML(body.String(), c.CallName())
		body = bytes.NewBufferString(bodyStr)
	}

	if e.Logger != nil {
		e.Logger(body.String())
	}

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/ws/api.dll", e.baseUrl), body)
	req.Header.Add("X-EBAY-API-DEV-NAME", e.DevId)
	req.Header.Add("X-EBAY-API-APP-NAME", e.AppId)
	req.Header.Add("X-EBAY-API-CERT-NAME", e.CertId)
	req.Header.Add("X-EBAY-API-CALL-NAME", c.CallName())
	req.Header.Add("X-EBAY-API-SITEID", strconv.Itoa(e.SiteId))
	req.Header.Add("X-EBAY-API-COMPATIBILITY-LEVEL", strconv.Itoa(837))
	req.Header.Add("Content-Type", "text/xml")

	client := &http.Client{
		// Transport: &http.Transport{
		// 	Proxy: func(_ *http.Request) (*url.URL, error) {
		// 		return url.Parse("http://127.0.0.1:8888")
		// 	},
		// },
	}
	resp, err := client.Do(req)
	if err != nil {
		if urlErr, ok := err.(*url.Error); ok {
			return OtherEbayResponse{}, urlErr
		}
		return OtherEbayResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		httpErr := httpError{statusCode: resp.StatusCode}
		httpErr.body, _ = io.ReadAll(resp.Body)
		return OtherEbayResponse{}, httpErr
	}

	bodyContents, _ := io.ReadAll(resp.Body)
	if e.Logger != nil {
		e.Logger(string(bodyContents))
	}

	response, err := c.ParseResponse(bodyContents)
	if response.Failure() {
		return response, response.ResponseErrors()
	}

	return response, err
}
