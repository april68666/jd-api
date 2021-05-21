package jdapi

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"time"

	"github.com/guonaihong/gout"
)

type Api struct {
	appKey    string
	appSecret string
}

func NewJdApi(appKey, appSecret string) *Api {
	return &Api{
		appKey:    appKey,
		appSecret: appSecret,
	}
}

// PositionCreate 创建推广位
func (a *Api) PositionCreate(param PositionCreateParams) (*PositionCreateResp, error) {
	resp := PositionCreateResp{}
	bytes, err := Marshal(param)
	if err != nil {
		return nil, err
	}
	bodyStr := string(bytes)
	currentTime := strconv.FormatInt(time.Now().Unix(), 10)
	singBody := a.appKey + a.appSecret + currentTime + bodyStr
	sign := GetMd5Encode(singBody)
	url := "https://jyapi.jd.com/open-api/jy/position/create?appKey=" + a.appKey + "&sign=" + sign + "&current=" + currentTime
	err = gout.POST(url).Debug(true).SetJSON(&param).BindJSON(&resp).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PositionQuery 查询推广位
func (a *Api) PositionQuery(param PositionQueryParams) (*PositionQueryResp, error) {
	resp := PositionQueryResp{}
	bytes, err := Marshal(param)
	if err != nil {
		return nil, err
	}
	bodyStr := string(bytes)
	currentTime := strconv.FormatInt(time.Now().Unix(), 10)
	singBody := a.appKey + a.appSecret + currentTime + bodyStr
	sign := GetMd5Encode(singBody)
	url := "https://jyapi.jd.com/open-api/jy/position/query?appKey=" + a.appKey + "&sign=" + sign + "&current=" + currentTime
	err = gout.POST(url).SetJSON(&param).BindJSON(&resp).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GoodsQuery 商品查询
func (a *Api) GoodsQuery(param GoodsQueryParams) (*GoodsQueryResp, error) {
	resp := GoodsQueryResp{}
	bytes, err := Marshal(param)
	if err != nil {
		return nil, err
	}
	bodyStr := string(bytes)
	currentTime := strconv.FormatInt(time.Now().Unix(), 10)
	singBody := a.appKey + a.appSecret + currentTime + bodyStr
	sign := GetMd5Encode(singBody)
	url := "https://jyapi.jd.com/open-api/jy/goods/query?appKey=" + a.appKey + "&sign=" + sign + "&current=" + currentTime
	err = gout.POST(url).SetJSON(&param).BindJSON(&resp).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PromotionGet 获取推广链接
func (a *Api) PromotionGet(param PromotionGetParams) (*PromotionGetResp, error) {
	resp := PromotionGetResp{}
	bytes, err := Marshal(param)
	if err != nil {
		return nil, err
	}
	bodyStr := string(bytes)
	currentTime := strconv.FormatInt(time.Now().Unix(), 10)
	singBody := a.appKey + a.appSecret + currentTime + bodyStr
	sign := GetMd5Encode(singBody)
	url := "https://jyapi.jd.com/open-api/jy/promotion/get?appKey=" + a.appKey + "&sign=" + sign + "&current=" + currentTime
	err = gout.POST(url).SetJSON(&param).BindJSON(&resp).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// OrderQuery 订单查询
func (a *Api) OrderQuery(param OrderQueryParams) (*OrderQueryResp, error) {
	resp := OrderQueryResp{}
	bytes, err := Marshal(param)
	if err != nil {
		return nil, err
	}
	bodyStr := string(bytes)
	currentTime := strconv.FormatInt(time.Now().Unix(), 10)
	singBody := a.appKey + a.appSecret + currentTime + bodyStr
	sign := GetMd5Encode(singBody)
	url := "https://jyapi.jd.com/open-api/jy/order/query-items?appKey=" + a.appKey + "&sign=" + sign + "&current=" + currentTime
	err = gout.POST(url).SetJSON(&param).BindJSON(&resp).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func GetMd5Encode(data string) string {
	h := md5.New()
	_, _ = h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
