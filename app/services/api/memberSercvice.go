package services

import (
	"dqh-test/libs/curl"
	"encoding/json"
	"errors"
	"time"
)

type Wallet struct {
	Status int `json:"status"`
	Data  Data `json:"data"`
}

type Data struct {
	Address string `json:"address"`
	PrivateKey string `json:"privateKey"`
}

//获取用户充币地址OMNI
func GetTokenAddress() ( *Wallet,  error){
	var wallet Wallet
	url := "http://182.16.48.170:8989/api/account/xdpool/btc/createwallet?access_key=xdpool"
	requset := curl.Post(url).DialTimeout(time.Second* 30)
	response,err := requset.Do()
	if err != nil{
		return &wallet,err
	}
	//var json_str = "{\"status\":0,\"data\":{\"address\":\"1Lsyqf2x1eGgR4mTrETs4hJ5MWsotTrJQW\",\"privateKey\":\"TJ9SvUcrj+OQ2+d/5rPugB5LwU1FxF4M4FxL+UvLsRwXDCpBAwJdbf0Xjt1MetvWlAxU69s+Cfpe5JizTTmvEA==\"}}"
	//response := curl.Response{
	//	StatusCode :200,
	//	Body : json_str,
	//}
	if response.StatusCode != 200  || len(response.Body) == 0{
		return &wallet, errors.New("请求错误")
	}
	json.Unmarshal([]byte(response.Body), &wallet)
	return &wallet,nil
}

//获取用户充币地址ERC20
func Geterc20Address()(wallte *Wallet, err error)  {
	url := "http://182.16.48.170:8989/api/account/xdpool/createwallet?access_key=xdpool"
	requset := curl.Post(url).DialTimeout(time.Second* 30)
	response,err := requset.Do()
	if err != nil{
		return wallte,err
	}
	if response.StatusCode != 200  || len(response.Body) == 0{
		return wallte, errors.New("请求错误")
	}
	json.Unmarshal([]byte(response.Body), wallte)
	return wallte,nil
}
