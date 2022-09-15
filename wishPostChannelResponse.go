package wishPost

type WishPostPickupService struct {
	Pickup_service_code string `json:"pickup_service_code"`
	Pickup_service_name string `json:"pickup_service_name"`
}

type WishPostChannel struct {
	Pickup_services        []WishPostPickupService `json:"pickup_services"`
	Is_registered          bool                    `json:"is_registered"`
	Support_warehouses     bool                    `json:"support_warehouses"`
	Is_only_online_payment bool                    `json:"is_only_online_payment"`
	Channel_name           string                  `json:"channel_name"`
	Otype                  string                  `json:"otype"`
	Allow_battery          bool                    `json:"allow_battery"`
	Channel_name_en        string                  `json:"channel_name_en"`
	Warehouses             []string                `json:"warehouses"`
	Sensitivity_types      []int                   `json:"sensitivity_types"`
}

type WishPostChannelResponse struct {
	BaseResponse
	Data []WishPostChannel `json:"data"`
}
