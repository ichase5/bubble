package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
)

func main() {

	//连接数据库
	err := dao.InitMySQL()
	defer dao.Close()//退出main函数时执行
	if err!=nil{
		//抛出异常
		panic(err)
	}

	//自动创建表 (表名todos)
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetUpRouter()

	r.Run(":80")

}
