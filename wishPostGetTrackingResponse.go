package wishPost

import "encoding/xml"

type WishPostTrackInfo struct {
	Date          string `xml:"date"`
	Status_number string `xml:"status_Number"`
	Status_desc   string `xml:"status_Desc"`
	Remark        string `xml:"remark"`
	City          string `xml:"city"`
	State         string `xml:"state"`
	Country_code  string `xml:"country_Code"`
}

type WishPostTracks struct {
	XMLName       xml.Name            `xml:"tracks"`
	Barcode       string              `xml:"barcode,attr"`
	Error_message string              `xml:"error_message,attr"`
	Track         []WishPostTrackInfo `xml:"track"`
}

type WishPostGetTrackingResponse struct {
	XMLName   xml.Name         `xml:"root"`
	Status    int              `xml:"status"`
	Timestamp string           `xml:"timestamp"`
	Tracks    []WishPostTracks `xml:"tracks"`
}
