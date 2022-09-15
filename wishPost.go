package wishPost

import (
	"encoding/json"
	"encoding/xml"
	"github.com/wms3001/goCommon"
	"github.com/wms3001/http"
	"github.com/wms3001/logisticsCommon"
)

var url string = "https://www.wishpost.cn"

type WishPost struct {
	Accesstoken string `json:"accesstoken"`
	ContentType string `json:"contentType"`
}

/**
测试token
*/
func (wishPost *WishPost) TestToken() *goCommon.Resp {
	m := map[string]string{
		"authorization": "Bearer " + wishPost.Accesstoken,
		"Content-Type":  wishPost.ContentType,
	}
	a := map[string]string{
		"Accept": "application/json",
		//"Accept": "text/plain",
	}
	var wHttp = http.Http{
		url + "/api/v3/oauth/test",
		"",
		"",
		m,
		a,
		"",
		10,
		"",
		[]byte{},
	}
	resp := wHttp.Get()
	wishPostTokenTestResponse := &WishPostTokenTestResponse{}
	json.Unmarshal([]byte(resp.Data), wishPostTokenTestResponse)
	if wishPostTokenTestResponse.Code != 0 {
		resp.Message = wishPostTokenTestResponse.Message
		resp.HttpCode = wishPostTokenTestResponse.Code
	}
	return resp
}

/**
获取渠道
*/
func (wishPost *WishPost) GetChannels() *goCommon.Resp {
	m := map[string]string{
		"authorization": "Bearer " + wishPost.Accesstoken,
		"Content-Type":  wishPost.ContentType,
	}
	a := map[string]string{
		"Accept": "application/json",
		//"Accept": "text/plain",
	}
	var wHttp = http.Http{
		url + "/api/v3/channels",
		"",
		"",
		m,
		a,
		"",
		10,
		"",
		[]byte{},
	}
	resp := wHttp.Get()
	wishPostChannelResponse := &WishPostChannelResponse{}
	json.Unmarshal([]byte(resp.Data), wishPostChannelResponse)
	if wishPostChannelResponse.Code != 0 {
		resp.Message = wishPostChannelResponse.Message
		resp.HttpCode = wishPostChannelResponse.Code
	}
	var channles []logisticsCommon.LogisticsChannel
	for _, v := range wishPostChannelResponse.Data {
		var channel logisticsCommon.LogisticsChannel
		channel.ChannelNameCn = v.Channel_name
		channel.ChannelNameEn = v.Channel_name_en
		channel.ChannelCode = v.Otype
		channles = append(channles, channel)
	}
	res, _ := json.Marshal(channles)
	resp.Data = string(res)
	return resp
}

/**
获取账号信息
*/
func (wishPost *WishPost) GetAccount() *goCommon.Resp {
	m := map[string]string{
		"authorization": "Bearer " + wishPost.Accesstoken,
		"Content-Type":  wishPost.ContentType,
	}
	a := map[string]string{
		"Accept": "application/json",
		//"Accept": "text/plain",
	}
	var wHttp = http.Http{
		url + "/api/v3/account_info",
		"",
		"",
		m,
		a,
		"",
		10,
		"",
		[]byte{},
	}
	resp := wHttp.Get()
	wishPostAccountResponse := &WishPostAccountResponse{}
	json.Unmarshal([]byte(resp.Data), wishPostAccountResponse)
	if wishPostAccountResponse.Code != 0 {
		resp.Message = wishPostAccountResponse.Message
		resp.HttpCode = wishPostAccountResponse.Code
	}
	res, _ := json.Marshal(wishPostAccountResponse)
	resp.Data = string(res)
	return resp
}

/**
获取授权范围
*/
func (wishPost *WishPost) GetScopes() *goCommon.Resp {
	m := map[string]string{
		"authorization": "Bearer " + wishPost.Accesstoken,
		"Content-Type":  wishPost.ContentType,
	}
	a := map[string]string{
		"Accept": "application/json",
		//"Accept": "text/plain",
	}
	var wHttp = http.Http{
		url + "/api/v3/scopes",
		"",
		"",
		m,
		a,
		"{\"scope_names\": []\n}",
		10,
		"",
		[]byte{},
	}
	resp := wHttp.Post()
	wishPostScopesresponse := &WishPostScopesresponse{}
	json.Unmarshal([]byte(resp.Data), wishPostScopesresponse)
	if wishPostScopesresponse.Code != 0 {
		resp.Message = wishPostScopesresponse.Message
		resp.HttpCode = wishPostScopesresponse.Code
	}
	res, _ := json.Marshal(wishPostScopesresponse)
	resp.Data = string(res)
	return resp
}

/**
获取通知公告
*/
func (wishPost *WishPost) GetAnnouncements(wishPostAnnouncementRequest *WishPostAnnouncementRequest) *goCommon.Resp {
	m := map[string]string{
		"authorization": "Bearer " + wishPost.Accesstoken,
		"Content-Type":  wishPost.ContentType,
	}
	a := map[string]string{
		"Accept": "application/json",
		//"Accept": "text/plain",
	}
	param, _ := json.Marshal(wishPostAnnouncementRequest)
	req := string(param)
	var wHttp = http.Http{
		url + "/api/v3/announcements",
		"",
		"",
		m,
		a,
		req,
		10,
		"",
		[]byte{},
	}
	resp := wHttp.Post()
	wishPostAnnouncementResponse := &WishPostAnnouncementResponse{}
	json.Unmarshal([]byte(resp.Data), wishPostAnnouncementResponse)
	if wishPostAnnouncementResponse.Code != 0 {
		resp.Message = wishPostAnnouncementResponse.Message
		resp.HttpCode = wishPostAnnouncementResponse.Code
	}
	res, _ := json.Marshal(wishPostAnnouncementResponse)
	resp.Data = string(res)
	return resp
}

/**
获取跟踪信息
*/
func (wishPost *WishPost) GetTracking(wishPostGetTrackingRequest *WishPostGetTrackingRequest) *goCommon.Resp {
	m := map[string]string{
		"authorization": "Bearer " + wishPost.Accesstoken,
		"Content-Type":  wishPost.ContentType,
	}
	a := map[string]string{
		"Accept": "application/xml",
		//"Accept": "text/plain",
	}
	output, err := xml.MarshalIndent(wishPostGetTrackingRequest, " ", " ")
	if err != nil {
		resp := &goCommon.Resp{}
		resp.Code = -1
		resp.Message = err.Error()
		return resp
	}
	req := string(output)
	var wHttp = http.Http{
		url + "/api/v2/tracking",
		"",
		"",
		m,
		a,
		req,
		10,
		"",
		[]byte{},
	}
	resp := wHttp.Post()
	wishPostGetTrackingResponse := &WishPostGetTrackingResponse{}
	xml.Unmarshal([]byte(resp.Data), wishPostGetTrackingResponse)
	res, _ := json.Marshal(wishPostGetTrackingResponse)
	resp.Data = string(res)
	return resp
}

func (wishPost *WishPost) GetWarehouse(request *WishPostWarehouseRequest) *goCommon.Resp {
	m := map[string]string{
		"authorization": "Bearer " + wishPost.Accesstoken,
		"Content-Type":  wishPost.ContentType,
	}
	a := map[string]string{
		"Accept": "application/json",
		//"Accept": "text/plain",
	}
	param, _ := json.Marshal(request)
	req := string(param)
	var wHttp = http.Http{
		url + "/api/v3/warehouse",
		"",
		"",
		m,
		a,
		req,
		10,
		"",
		[]byte{},
	}
	resp := wHttp.Post()
	wishPostAnnouncementResponse := &WishPostAnnouncementResponse{}
	json.Unmarshal([]byte(resp.Data), wishPostAnnouncementResponse)
	if wishPostAnnouncementResponse.Code != 0 {
		resp.Message = wishPostAnnouncementResponse.Message
		resp.HttpCode = wishPostAnnouncementResponse.Code
	}
	res, _ := json.Marshal(wishPostAnnouncementResponse)
	resp.Data = string(res)
	return resp
}
