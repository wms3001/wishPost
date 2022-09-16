# golang wishPost api对接
## 简介
实现对接wish邮api接口
## 使用
```
go get github.com/wms3001/wishPost
```
## 实例
1. 测试token
```go
var wishPost = &WishPost{}
wishPost.Accesstoken = ""
wishPost.ContentType = "application/json"
resp := wishPost.TestToken()
```
2. 获取渠道
```go
var wishPost = &WishPost{}
wishPost.Accesstoken = ""
wishPost.ContentType = "application/json"
resp := wishPost.GetChannels()
```
3. 获取账号信息
```go
var wishPost = &WishPost{}
wishPost.Accesstoken = ""
wishPost.ContentType = "application/json"
resp := wishPost.GetAccount()
```
4. 获取授权范围
```go
var wishPost = &WishPost{}
wishPost.Accesstoken = ""
wishPost.ContentType = "application/json"
resp := wishPost.GetScopes()
```
5. 获取账号通知
```go
var wishPost = &WishPost{}
wishPost.Accesstoken = ""
wishPost.ContentType = "application/json"
var wishPostAnnouncementRequest = &WishPostAnnouncementRequest{}
wishPostAnnouncementRequest.Start_date = "2022-08-01"
wishPostAnnouncementRequest.End_date = "2022-08-20"
wishPostAnnouncementRequest.Offset = 0
wishPostAnnouncementRequest.Limit = 50
resp := wishPost.GetAnnouncements(wishPostAnnouncementRequest)
```
6. 获取跟踪信息
```go
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
```
7. 获取仓库信息
```go
var wishPost = &WishPost{}
wishPost.Accesstoken = ""
wishPost.ContentType = "application/json"
wishPostWarehouseRequest := &WishPostWarehouseRequest{}
//wishPostWarehouseRequest.Warehouse_ids = append(wishPostWarehouseRequest.Warehouse_ids, 1001)
wishPostWarehouseRequest.Warehouse_ids = []int{1001, 1002, 1003}
resp := wishPost.GetWarehouse(wishPostWarehouseRequest)
```
8. 取消订单
```go
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
```
9. 获取订单
```go
var wishPost = &WishPost{}
wishPost.Accesstoken = ""
wishPost.ContentType = "application/json"
wishPostOrderStatusRequest := &WishPostOrderStatusRequest{}
wishPostOrderStatusRequest.Access_token = ""
wishPostOrderStatusRequest.Wish_standard_tracking_ids = append(wishPostOrderStatusRequest.Wish_standard_tracking_ids, "ewr234234234")
resp := wishPost.GetOrder(wishPostOrderStatusRequest)
```
10. 创建订单
```go
var wishPost = &WishPost{}
wishPost.Accesstoken = ""
wishPost.ContentType = "application/xml"
wishPostCreateOrderRequest := &WishPostCreateOrderRequest{}
resp := wishPost.CreateOrder(wishPostCreateOrderRequest)
```