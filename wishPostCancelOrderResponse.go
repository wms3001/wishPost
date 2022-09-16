package wishPost

type WishPostCancelOrderData struct {
	Cancel_result_code        int    `json:"cancel_result_code"`
	Wish_standard_tracking_id string `json:"wish_standard_tracking_id"`
	Cancel_result_message     string `json:"cancel_result_message"`
	User_id                   string `json:"user_id"`
	Tracking_id               string `json:"tracking_id"`
}

type WishPostCancelOrderResponse struct {
	BaseResponse
	Failed_cancel_orders []WishPostCancelOrderData `json:"failed_cancel_orders"`
}
