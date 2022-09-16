package wishPost

import "encoding/xml"

type WishPostCreateOrderResponse struct {
	XMLName                   xml.Name `xml:"root"`
	Status                    int      `xml:"status"`
	Timestamp                 string   `xml:"timestamp"`
	Error_message             string   `xml:"error_message"`
	Guid                      string   `xml:"guid,attr"`
	Error_ref_id              string   `xml:"error_ref_id,attr"`
	Mark                      string   `xml:"mark"`
	PDF_10_EN_URL             string   `xml:"PDF_10_EN_URL"`
	PDF_10_LCL_URL            string   `xml:"PDF_10_LCL_URL"`
	Wish_standard_tracking_id string   `xml:"wish_standard_tracking_id"`
	State                     string   `xml:"state"`
	Barcode                   string   `xml:"barcode"`
}
