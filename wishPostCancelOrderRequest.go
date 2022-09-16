package wishPost

type WishPostCancelOrder struct {
	Tracking_id        string `json:"tracking_id"`
	Cancel_reason_code string `json:"cancel_reason_code"`
}

type WishPostCancelOrderRequest struct {
	Access_token string                `json:"access_token"`
	Orders       []WishPostCancelOrder `json:"orders"`
}
