package controller

import (
	"Bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK,"index.html",nil)
}

func CreateATodo(c *gin.Context) {
	//前端页面填写待办事项 点击提交，会发送到请求这里
	//1.从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	//2.存入数据库
	err := models.CreateATodo(&todo)
	//3.返回响应
	if err!=nil {//错误
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{//正常
		c.JSON(http.StatusOK,todo)
	}

}

func GetTodoList(c *gin.Context) {
	//var todoList []models.Todo
	todoList,err := models.GetAllTodo()
	if err!=nil {
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusOK,todoList)
	}

}

func UpdateATodo(c *gin.Context) {
	id,ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK,gin.H{"error":"无效的id"})
		return
	}
	//var todo models.Todo
	todo,err := models.GetATodo(id)
	if err!=nil {
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err =models.UpdateATodo(todo);err!=nil {
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusOK,todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id,ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK,gin.H{"error":"无效的id"})
		return
	}
	if err := models.DeleteATodo(id);err!=nil {
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusOK,gin.H{"id":"deleted"})
	}
}