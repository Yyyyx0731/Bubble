package routers

import (
	"Bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine{
	r := gin.Default()

	//告诉gin框架 模板文件引用的静态文件去哪里找
	r.Static("/static","static")
	//告诉gin框架去那六找模板文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.IndexHandler)

	//定义应该api v1
	v1Group := r.Group("/v1")
	{
		//待办事项

		//添加
		v1Group.POST("/todo",controller.CreateATodo)

		//查看所有待办事项
		v1Group.GET("/todo", controller.GetTodoList)

		//修改某一个待办事项
		v1Group.PUT("/todo/:id",controller.UpdateATodo)

		//删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}

	return r
}
