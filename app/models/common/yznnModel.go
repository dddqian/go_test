package models

import (
	mysql "dqh-test/libs/database/mysql"
)

type Yznn struct {
	Id int `json:"id"`
	Uid int `json:"uid"`
	UserId string `json:"user_id"`
	YzCode string `json:"yz_code"`
	YzAddress string `json:"yz_address"`
	YzType int `json:"yz_type"`
	STime int64 `json:"s_time"`
	ETime int64 `json:"e_time"`
	Info string `json:"info"`
}

func (y *Yznn) TableName() string {
	return "sf_yznn"
}

func GetYznnInfo(yznn *Yznn) (error)  {
	mysql := mysql.GetMysqlInstance()
	if mysql.Error != nil{
		return  mysql.Error
	}
	result := mysql.DB.Where(yznn).First(yznn)
	return  result.Error
}
