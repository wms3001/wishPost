package wishPost

type WishPostOrderStatusRequest struct {
	Access_token               string   `json:"access_token"`
	Wish_standard_tracking_ids []string `json:"wish_standard_tracking_ids"`
}
