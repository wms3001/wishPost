package wishPost

import (
	"fmt"
	"testing"
)

func TestWishPost_TestToken(t *testing.T) {
	var wishPost = &WishPost{}
	wishPost.Accesstoken = "fa486cd391724eb9a459147158ee3e12"
	wishPost.ContentType = "application/json"
	resp := wishPost.TestToken()
	println(resp.Code)
	println(resp.Message)
	println(resp.Data)
}

func TestWishPost_GetChannels(t *testing.T) {
	var wishPost = &WishPost{}
	wishPost.Accesstoken = "fa486cd391724eb9a459147158ee3e12"
	wishPost.ContentType = "application/json"
	resp := wishPost.GetChannels()
	fmt.Println(resp.Code)
	fmt.Println(resp.Message)
	fmt.Println(resp.Data)
}

func TestWishPost_GetAccount(t *testing.T) {
	var wishPost = &WishPost{}
	wishPost.Accesstoken = "fa486cd391724eb9a459147158ee3e12"
	wishPost.ContentType = "application/json"
	resp := wishPost.GetAccount()
	fmt.Println(resp.Code)
	fmt.Println(resp.Message)
	fmt.Println(resp.Data)
}

func TestWishPost_GetScopes(t *testing.T) {
	var wishPost = &WishPost{}
	wishPost.Accesstoken = "fa486cd391724eb9a459147158ee3e12"
	wishPost.ContentType = "application/json"
	resp := wishPost.GetScopes()
	fmt.Println(resp.Code)
	fmt.Println(resp.Message)
	fmt.Println(resp.Data)
}

func TestWishPost_GetAnnouncements(t *testing.T) {
	var wishPost = &WishPost{}
	wishPost.Accesstoken = "fa486cd391724eb9a459147158ee3e12"
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
	wishPost.Accesstoken = "fa486cd391724eb9a459147158ee3e12"
	wishPost.ContentType = "application/xml"
	var wishPostGetTrackingRequest = &WishPostGetTrackingRequest{}
	wishPostGetTrackingRequest.Language = "cn"
	wishPostGetTrackingRequest.Access_token = "fa486cd391724eb9a459147158ee3e12"
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
	wishPost.Accesstoken = "fa486cd391724eb9a459147158ee3e12"
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
