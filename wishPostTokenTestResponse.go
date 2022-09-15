package wishPost

type WishPostUserId struct {
	User_id string `json:"user_id"`
}

type WishPostTokenTestResponse struct {
	BaseResponse
	Data WishPostUserId `json:"data"`
}
