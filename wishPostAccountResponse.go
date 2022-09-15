package wishPost

type WishPostAccount struct {
	Total          string `json:"total"`
	Unavailable    string `json:"unavailable"`
	Debt_frozen    string `json:"debt_frozen"`
	Debt_remaining string `json:"debt_remaining"`
	Debt_limit     string `json:"debt_limit"`
	Currency       string `json:"currency"`
	Withdrawable   string `json:"withdrawable"`
	Debt           string `json:"debt"`
	Remaining      string `json:"remaining"`
}

type WishPostAccountResponse struct {
	BaseResponse
	Data WishPostAccount `json:"data"`
}
