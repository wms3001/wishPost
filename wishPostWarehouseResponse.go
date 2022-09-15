package wishPost

type WishPostWarehouseInfo struct {
	Province     string `json:"province"`
	City         string `json:"city"`
	Name         string `json:"name"`
	District     string `json:"district"`
	Zipcode      string `json:"zipcode"`
	Address      string `json:"address"`
	Phone_number string `json:"phone_number"`
	Email        string `json:"email"`
}

type WishPostWarehouse struct {
	Official_name_en     string                `json:"official_name_en"`
	Official_name_cn     string                `json:"official_name_cn"`
	Enum_value           int                   `json:"enum_value"`
	Warehouse_address_cn WishPostWarehouseInfo `json:"warehouse_address_cn"`
	Warehouse_address_en WishPostWarehouseInfo `json:"warehouse_address_en"`
}

type WishPostWarehouseResponse struct {
	BaseResponse
	Warehouses []WishPostWarehouse `json:"warehouses"`
}
