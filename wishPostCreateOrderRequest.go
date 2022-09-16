package wishPost

import "encoding/xml"

type WishPostProductItem struct {
	XMLName                       xml.Name `xml:"product_items"`
	Wish_order_id                 string   `xml:"wish_order_id"`
	Product_item_from_country     string   `xml:"product_item_from_country"`
	Product_item_name             string   `xml:"product_item_name"`
	Product_item_name_chinese     string   `xml:"product_item_name_chinese"`
	Product_item_num              int      `xml:"product_item_num"`
	Product_item_weight           float64  `xml:"product_item_weight"`
	Product_item_single_price     float64  `xml:"product_item_single_price"`
	Product_item_has_battery      int      `xml:"product_item_has_battery"`
	Product_item_sensitivity_type int      `xml:"product_item_sensitivity_type"`
	Product_item_product_url      string   `xml:"product_item_product_url"`
}

type WishPostOrder struct {
	Guid                     string                `xml:"guid"`
	Otype                    string                `xml:"otype"`
	Warehouse_code           int                   `xml:"warehouse_code"`
	Pickup_from_local        string                `xml:"pickup_from_local"`
	Pickup_addres_local      string                `xml:"pickup_addres_local"`
	Pickup_province_local    string                `xml:"pickup_province_local"`
	Pickup_city_local        string                `xml:"pickup_city_local"`
	Pickup_phone             string                `xml:"pickup_phone"`
	From                     string                `xml:"from"`
	Sender_addres            string                `xml:"sender_addres"`
	Sender_province          string                `xml:"sender_province"`
	Sender_city              string                `xml:"sender_city"`
	Sender_phone             string                `xml:"sender_phone"`
	Sender_zipcode           string                `xml:"sender_zipcode"`
	To                       string                `xml:"to"`
	To_local                 string                `xml:"to_local"`
	Recipient_addres         string                `xml:"recipient_addres"`
	Recipient_addres_local   string                `xml:"recipient_addres_local"`
	Recipient_country        string                `xml:"recipient_country"`
	Recipient_country_short  string                `xml:"recipient_country_short"`
	Recipient_country_local  string                `xml:"recipient_country_local"`
	Recipient_province       string                `xml:"recipient_province"`
	Recipient_province_local string                `xml:"recipient_province_local"`
	Recipient_city           string                `xml:"recipient_city"`
	Recipient_city_local     string                `xml:"recipient_city_local"`
	Recipient_postcode       string                `xml:"recipient_postcode"`
	Recipient_phone          string                `xml:"recipient_phone"`
	Type_no                  int                   `xml:"type_no"`
	Unit_measurement         string                `xml:"unit_measurement"`
	Unit_measurement_chinese string                `xml:"unit_measurement_chinese"`
	User_desc                string                `xml:"user_desc"`
	Trade_amount             float64               `xml:"trade_amount"`
	Receive_from             string                `xml:"receive_from"`
	Receive_province         string                `xml:"receive_province"`
	Receive_city             string                `xml:"receive_city"`
	Receive_addres           string                `xml:"receive_addres"`
	Receive_phone            string                `xml:"receive_phone"`
	Doorpickup               int                   `xml:"doorpickup"`
	product_items            []WishPostProductItem `xml:"product_items"`
}

type WishPostCreateOrderRequest struct {
	XMLName      xml.Name        `xml:"orders"`
	Access_token string          `xml:"access_token"`
	mark         string          `xml:"mark"`
	bid          int             `xml:"bid"`
	Order        []WishPostOrder `xml:"order"`
}
