package controller

import (
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context){
	c.HTML(http.StatusOK,"index.html",nil)
}

func CreateHandler(c *gin.Context){
	var todo models.Todo
	c.BindJSON(&todo)//json解析，填充todo的title字段
	err := models.CreateTodo(&todo)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"error" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,todo)
}

func ViewAllHandler(c *gin.Context){
	todoList, err := models.GetAllTodo()
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"error" : err.Error(),
		})
	} else{
		c.JSON(http.StatusOK,todoList)
	}
}

func UpdateHandler(c *gin.Context){
	id,ok := c.Params.Get("id")
	if !ok{
		c.JSON(http.StatusOK,gin.H{
			"error" : "id错误",
		})
		return
	}
	todo, err := models.GetTheTodo(id)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"error" : err.Error(),
		})
		return
	}
	//修改为哪种状态(json参数)
	c.BindJSON(todo)
	if err=models.UpdateTodo(todo);err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"error" : err.Error(),
		})
	}else{
		c.JSON(http.StatusOK,todo)
	}
}


func DeleteHandler(c *gin.Context){
	id,ok := c.Params.Get("id")
	if !ok{
		c.JSON(http.StatusOK,gin.H{
			"error" : "id错误",
		})
		return
	}
	if err:=models.DeleteTodo(id);err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"error" : err.Error(),
		})
	} else{
		c.JSON(http.StatusOK,gin.H{
			id : "deleted",
		})
	}
}



