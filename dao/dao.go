package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //只执行init，不导入包
)

//全局变量
var (
	DB *gorm.DB
)

func InitMySQL()(err error){
	dsn := "root:Admin123!@tcp(121.4.35.23:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	//注意，不能用 :=
	DB, err = gorm.Open("mysql", dsn)
	if err!=nil{
		return
	}
	//ping检测
	return DB.DB().Ping()
}

func Close(){
	DB.Close()
}

