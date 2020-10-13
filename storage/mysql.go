package storage

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopackage/loger"
)

var DB *gorm.DB

/**
	docker容器内访问宿主机MySql，修改监听地址为0.0.0.0  /etc/mysql/mysql.conf.d/mysqld.cnf
 */
func NewStorageConnetcion(dbType string, dbHost string, dbName string, dbUser string, dbPassword string) *gorm.DB{
	var err error

	if dbType == "mysql" {
		DB, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			dbUser,
			dbPassword,
			dbHost,
			dbName))
	}else if dbType == "postgres"{

		DB, err = gorm.Open(dbType, fmt.Sprintf("host=%s dbname=%s user=%s sslmode=disable password=%s",
			dbHost,
			dbName,
			dbUser,
			dbPassword))
	}

	if err != nil {
		panic(dbType + " Connect Failed !" + err.Error())
	}

	DB.LogMode(true)
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	DB.SetLogger(loger.Loger)

	loger.Info(dbType + " Connect Success.", dbHost)
	return DB
}


