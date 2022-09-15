package wishPost

type WishPostScope struct {
	Resource     string   `json:"resource"`
	Name         string   `json:"name"`
	Apis         []string `json:"apis"`
	Chinese_name string   `json:"chinese_name"`
	Role         string   `json:"role"`
	English_name string   `json:"english_name"`
	Attribute    string   `json:"attribute"`
	Action       string   `json:"action"`
}

type WishPostScopesresponse struct {
	BaseResponse
	Scopes []WishPostScope `json:"scopes"`
}
