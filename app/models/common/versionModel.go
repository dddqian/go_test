package models

import "dqh-test/libs/database/mysql"

type Version struct {
	Id int `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	AppType string `json:"app_type"`
	AppVersion string `json:"app_version"`
	AppLink string `json:"app_link"`
	Addtime int64 `json:"addtime"`
	AppContent string `json:"app_content"`
}

func (v *Version) TableName() string {
	return "sf_version"
}

func GetVersionInfo(app_type string,ver string) (*Version,error){
	var(
		version  Version
	)

	mysql := mysql.GetMysqlInstance()
	if mysql.Error != nil{
		return &version ,mysql.Error
	}

	result := mysql.DB.Where("app_type = ? and app_version = ?", app_type, ver).First(&version)

	return  &version ,result.Error
}

