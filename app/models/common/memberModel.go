package models

import "dqh-test/libs/database/mysql"

type Member struct {
	Id int `json:"id"`
 	Yaoqingma string `json:"yaoqingma"`
	Bianma string `json:"bianma"`
	Tel string `json:"tel"`
	Pwd string `json:"pwd"`
	ReName int `json:"re_name"`
	IdCard int `json:"id_card"`
	Status int `json:"status"`
	Addtime int64 `json:"addtime"`
	RzTime int64 `json:"rz_time"`
	Updatetime int64 `json:"updatetime"`
	Token string `json:"token"`
	IdcardZImg string `json:"idcard_z_img"`
	IdcardFImg string `json:"idcard_f_img"`
	JyPwd string `json:"jy_pwd"`
	ReId int `json:"re_id"`
	IsStatus int `json:"is_status"`
	RePath string `json:"re_path"`
	Level int `json:"level"`
	ReLevel int `json:"re_level"`
	Usdt float32 `json:"usdt"`
	EthAmount float32 `json:"eth_amount"`
	ZsSuanli float32 `json:"zs_suanli"`
	Zsl float32 `json:"zsl"`
	Address string `json:"address"` //BTC/USDT钱包地址
	PrivateKey string `json:"privateKey" gorm:"Column:privateKey"`
	ErcAddress string `json:"erc_address"`
	ErcPrivateKey string `json:"erc_privateKey" gorm:"Column:erc_privateKey"`
	TrcAddress string `json:"trc_address"`
	IntroCode string `json:"intro_code"` //邀请码
	Erweima string `json:"erweima" gorm:"default:'ceshi'"`

}

func (m *Member) TableName() string {
	return "sf_member"
}

func (m *Member) Insert() (error) {
	mysql := mysql.GetMysqlInstance()
	if mysql.Error != nil{
		return  mysql.Error
	}
	return  mysql.DB.Create(m).Error
}

func GetMenberInfoByTel(tel string) (*Member , error) {
	var (
		err error
		member Member
	)
	mysql := mysql.GetMysqlInstance()
	if mysql.Error != nil{
		return  &member,mysql.Error
	}
	result := mysql.DB.Where("tel=?", tel).First(&member)
	if result.Error != nil{
		err = result.Error
	}
	return &member,err
}

func UpdateMemberInfo(member *Member) (err error){
	mysql := mysql.GetMysqlInstance()
	if mysql.Error != nil{
		return  mysql.Error
	}
	result := mysql.DB.Model(member).Update("token")
	if result.Error != nil{
		err = result.Error
	}
	return
}

//根据邀请码获取用户信息
func GetMenberInfoByIntrocode(intro_code string)(*Member , error){
	var (
		err error
		member Member
	)
	mysql := mysql.GetMysqlInstance()
	if mysql.Error != nil{
		return  &member,mysql.Error
	}
	result := mysql.DB.Where("intro_code=?", intro_code).First(&member)
	if result.Error != nil{
		err = result.Error
	}
	return &member,err
}