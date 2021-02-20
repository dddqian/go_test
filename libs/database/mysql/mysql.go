package mysql

import (
	"dqh-test/libs/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"sync"
)


var mysqlManager *MysqlManager
var mysqlOnce sync.Once

func GetMysqlInstance() *MysqlManager{
	mysqlOnce.Do(func() {
		mysqlManager = new(MysqlManager)
		mysqlManager.initDB()
	})
	return mysqlManager
}


type MysqlManager struct {
	DB *gorm.DB
	Error error
}

func (m *MysqlManager) initDB() (*gorm.DB,error){
	dbconf, conf_err := config.GetConf("mysql", "default.master")
	if conf_err != nil {
		m.Error = conf_err
	}
	var db_err error
	m.DB, db_err = gorm.Open("mysql", dbconf["source"])

	if db_err != nil {
		m.Error = db_err
	}
	if m.DB.Error != nil {
		m.Error = m.DB.Error
	}
	m.DB.DB().SetMaxIdleConns(5)
	m.DB.SingularTable(true)

	return m.DB, m.Error
}

