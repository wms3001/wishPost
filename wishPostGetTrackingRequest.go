package wishPost

import "encoding/xml"

type WishPostTrack struct {
	Barcode string `xml:"barcode"`
}

type WishPostGetTrackingRequest struct {
	XMLName      xml.Name        `xml:"tracks"`
	Access_token string          `xml:"access_Token"`
	Track        []WishPostTrack `xml:"track"`
	Language     string          `xml:"language"`
}
