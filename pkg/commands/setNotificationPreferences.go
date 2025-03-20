package commands

import (
	"encoding/xml"

	"github.com/dracoDevs/go-ebay-plus/internal/ebay"
)

type DeliveryURLDetail struct {
	DeliveryURL     string
	DeliveryURLName string
	Status          string
}

type ApplicationDeliveryPreferences struct {
	AlertEmail         string
	AlertEnable        string
	ApplicationEnable  string
	ApplicationURL     string
	DeliveryURLDetails []DeliveryURLDetail
	DeviceType         string
	PayloadVersion     string
}

type EventProperty struct {
	EventType string
	Name      string
	Value     string
}

type UserData struct {
	ExternalUserData string
}

type NotificationEnable struct {
	EventEnable string
	EventType   string
}

type UserDeliveryPreferenceArray struct {
	NotificationEnable []NotificationEnable
}

type SetNotificationPreferences struct {
	ApplicationDeliveryPreferences ApplicationDeliveryPreferences
	DeliveryURLName                string
	EventProperty                  []EventProperty
	UserData                       UserData
	UserDeliveryPreferenceArray    UserDeliveryPreferenceArray
	ErrorLanguage                  string
	MessageID                      string
	Version                        string
	WarningLevel                   string
}

func (c SetNotificationPreferences) CallName() string {
	return "SetNotificationPreferences"
}

func (c SetNotificationPreferences) Body() interface{} {
	return c
}

func (c SetNotificationPreferences) ParseResponse(r []byte) (ebay.EbayResponse, error) {
	var xmlResponse SetNotificationPreferencesResponse
	err := xml.Unmarshal(r, &xmlResponse)

	return xmlResponse, err
}

type SetNotificationPreferencesResponse struct {
	ebay.OtherEbayResponse
}

func (r SetNotificationPreferencesResponse) ResponseErrors() ebay.EbayErrors {
	return r.OtherEbayResponse.Errors
}
