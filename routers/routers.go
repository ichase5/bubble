package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine{
	r := gin.Default()

	//去哪里找静态文件
	r.Static("/static","static")

	//去哪里找html文件
	r.LoadHTMLGlob("templates\\**")

	r.GET("/",controller.IndexHandler)

	group := r.Group("/v1")
	{
		//添加一个todo
		group.POST("/todo",controller.CreateHandler)

		//查看所有todo
		group.GET("/todo",controller.ViewAllHandler)

		//修改某个todo的状态,目标状态在json参数里
		group.PUT("/todo/:id",controller.UpdateHandler)

		group.DELETE("/todo/:id",controller.DeleteHandler)
	}
	return r
}





