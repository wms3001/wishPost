package wishPost

type WishPostStatusOrder struct {
	Wish_standard_tracking_id string                `json:"wish_standard_tracking_id"`
	Logistics_order_code      string                `json:"logistics_order_code"`
	Wish_transaction_ids      []string              `json:"wish_transaction_ids"`
	State_name                string                `json:"state_name"`
	State_code                int                   `json:"state_code"`
	Cancel_reason_code        string                `json:"cancel_reason_code"`
	Cancel_fail_reason        string                `json:"cancel_fail_reason"`
	Message                   string                `json:"message"`
	Reference_qrcode          string                `json:"reference_qrcode"`
	PDF_10_LCL_URL            string                `json:"PDF_10_LCL_URL"`
	PDF_10_EN_URL             string                `json:"PDF_10_EN_URL"`
	Dest_warehouse_code       int                   `json:"dest_warehouse_code"`
	Dest_warehouse_name       string                `json:"dest_warehouse_name"`
	Hub_warehouse_code        string                `json:"hub_warehouse_code"`
	Dest_warehouse_address_cn WishPostWarehouseInfo `json:"dest_warehouse_address_cn"`
}

type WishPostOrderStatusResponse struct {
	BaseResponse
	orders []WishPostStatusOrder `json:"orders"`
}
