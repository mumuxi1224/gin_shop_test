package models

import (
	"fmt"
	"gin_shop_test/pkg"
	"github.com/jinzhu/gorm"
	"log"
)

var (
	db *gorm.DB
)

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePerfix string
	)

	set, err := pkg.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("fail to get section 'database' %v ", err)
	}

	dbType = set.Key("TYPE").String()
	dbName = set.Key("DBNAME").String()
	user = set.Key("USER").String()
	password = set.Key("PASSWORD").String()
	host = set.Key("HOST").String()
	tablePerfix = set.Key("Table_Perfix").String()

	db, err = gorm.Open(dbType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user,
			password,
			host,
			dbName))
	if err != nil {
		panic(err)
	}

	// 启用Logger，显示详细日志
	db.LogMode(true)

	//连接池
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// 全局禁用表名复数
	db.SingularTable(true)

	//更改默认表名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePerfix + defaultTableName
	}
}

func closeDB() {
	defer db.Close()
}
