package commands

type Storefront struct {
	StoreCategoryID string
}

type ReturnPolicy struct {
	ReturnsAccepted, ReturnsAcceptedOption, ReturnsWithinOption, RefundOption, ShippingCostPaidByOption string
}

type ItemSpecifics struct {
	NameValueList []NameValueList
}

type NameValueList struct {
	Name  string
	Value []string
}

type PictureDetails struct {
	PictureURL string
}

type PrimaryCategory struct {
	CategoryID string
}

type BestOfferDetails struct {
	BestOfferEnabled bool
}

type BrandMPN struct {
	Brand, MPN string
}

type ProductListingDetails struct {
	UPC      string
	BrandMPN BrandMPN
}

type ShippingDetails struct {
	ShippingType                           string
	ShippingDiscountProfileID              string
	InternationalShippingDiscountProfileID string
	ShippingServiceOptions                 []ShippingServiceOption
	InternationalShippingServiceOption     []InternationalShippingServiceOption
}

type ShippingServiceOption struct {
	ShippingService               string
	ShippingServiceCost           float64
	ShippingServiceAdditionalCost float64
	FreeShipping                  bool
}

type InternationalShippingServiceOption struct {
	ShippingService               string
	ShippingServiceCost           float64
	ShippingServiceAdditionalCost float64
	ShipToLocation                []string
	ShippingServicePriority       int
}