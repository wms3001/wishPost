package wishPost

import (
	"fmt"
	"testing"
)

func TestWishPost_TestToken(t *testing.T) {
	var wishPost = &WishPost{}
	wishPost.Accesstoken = ""
	wishPost.ContentType = "application/json"
	resp := wishPost.TestToken()
	println(resp.Code)
	println(resp.Message)
	println(resp.Data)
}

func TestWishPost_GetChannels(t *testing.T) {
	var wishPost = &WishPost{}
	wishPost.Accesstoken = ""
	wishPost.ContentType = "application/json"
	resp := wishPost.GetChannels()
	fmt.Println(resp.Code)
	fmt.Println(resp.Message)
	fmt.Println(resp.Data)
}

func TestWishPost_GetAccount(t *testing.T) {
	var wishPost = &WishPost{}
	wishPost.Accesstoken = ""
	wishPost.ContentType = "application/json"
	resp := wishPost.GetAccount()
	fmt.Println(resp.Code)
	fmt.Println(resp.Message)
	fmt.Println(resp.Data)
}

func TestWishPost_GetScopes(t *testing.T) {
	var wishPost = &WishPost{}
	wishPost.Accesstoken = ""
	wishPost.ContentType = "application/json"
	resp := wishPost.GetScopes()
	fmt.Println(resp.Code)
	fmt.Println(resp.Message)
	fmt.Println(resp.Data)
}

func TestWishPost_GetAnnouncements(t *testing.T) {
	var wishPost = &WishPost{}
	wishPost.Accesstoken = ""
	wishPost.ContentType = "application/json"
	var wishPostAnnouncementRequest = &WishPostAnnouncementRequest{}
	wishPostAnnouncementRequest.Start_date = "2022-08-01"
	wishPostAnnouncementRequest.End_date = "2022-08-20"
	wishPostAnnouncementRequest.Offset = 0
	wishPostAnnouncementRequest.Limit = 50
	resp := wishPost.GetAnnouncements(wishPostAnnouncementRequest)
	fmt.Println(resp)
	fmt.Println(resp.Code)
	fmt.Println(resp.Message)
	fmt.Println(resp.Data)
}

func TestWishPost_GetTracking(t *testing.T) {
	var wishPost = &WishPost{}
	wishPost.Accesstoken = ""
	wishPost.ContentType = "application/xml"
	var wishPostGetTrackingRequest = &WishPostGetTrackingRequest{}
	wishPostGetTrackingRequest.Language = "cn"
	wishPostGetTrackingRequest.Access_token = ""
	wishPostTrack := WishPostTrack{}
	wishPostTrack.Barcode = "23wererwerwerwe"
	wishPostGetTrackingRequest.Track = append(wishPostGetTrackingRequest.Track, wishPostTrack)
	resp := wishPost.GetTracking(wishPostGetTrackingRequest)
	fmt.Println(resp)
	fmt.Println(resp.Code)
	fmt.Println(resp.Message)
	fmt.Println(resp.Data)
}

func TestWishPost_GetWarehouse(t *testing.T) {
	var wishPost = &WishPost{}
	wishPost.Accesstoken = ""
	wishPost.ContentType = "application/json"
	wishPostWarehouseRequest := &WishPostWarehouseRequest{}
	//wishPostWarehouseRequest.Warehouse_ids = append(wishPostWarehouseRequest.Warehouse_ids, 1001)
	wishPostWarehouseRequest.Warehouse_ids = []int{1001, 1002, 1003}
	resp := wishPost.GetWarehouse(wishPostWarehouseRequest)
	fmt.Println(resp)
	fmt.Println(resp.Code)
	fmt.Println(resp.Message)
	fmt.Println(resp.Data)
}

func TestWishPost_CancelOrders(t *testing.T) {
	var wishPost = &WishPost{}
	wishPost.Accesstoken = ""
	wishPost.ContentType = "application/json"
	wishPostCancelOrderRequest := &WishPostCancelOrderRequest{}
	wishPostCancelOrderRequest.Access_token = ""
	var wishPostCancelOrder WishPostCancelOrder
	wishPostCancelOrder.Tracking_id = "2342342342342"
	wishPostCancelOrder.Cancel_reason_code = "40101"
	wishPostCancelOrderRequest.Orders = append(wishPostCancelOrderRequest.Orders, wishPostCancelOrder)
	resp := wishPost.CancelOrders(wishPostCancelOrderRequest)
	fmt.Println(resp)
	fmt.Println(resp.Code)
	fmt.Println(resp.Message)
	fmt.Println(resp.Data)
}

func TestWishPost_GetOrder(t *testing.T) {
	var wishPost = &WishPost{}
	wishPost.Accesstoken = ""
	wishPost.ContentType = "application/json"
	wishPostOrderStatusRequest := &WishPostOrderStatusRequest{}
	wishPostOrderStatusRequest.Access_token = ""
	wishPostOrderStatusRequest.Wish_standard_tracking_ids = append(wishPostOrderStatusRequest.Wish_standard_tracking_ids, "ewr234234234")
	resp := wishPost.GetOrder(wishPostOrderStatusRequest)
	fmt.Println(resp)
	fmt.Println(resp.Code)
	fmt.Println(resp.Message)
	fmt.Println(resp.Data)
}

func TestWishPost_CreateOrder(t *testing.T) {
	var wishPost = &WishPost{}
	wishPost.Accesstoken = ""
	wishPost.ContentType = "application/xml"
	wishPostCreateOrderRequest := &WishPostCreateOrderRequest{}
	resp := wishPost.CreateOrder(wishPostCreateOrderRequest)
	fmt.Println(resp)
	fmt.Println(resp.Code)
	fmt.Println(resp.Message)
	fmt.Println(resp.Data)
}
