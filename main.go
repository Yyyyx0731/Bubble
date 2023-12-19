package main

import (
	"Bubble/dao"
	"Bubble/models"
	"Bubble/routers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main() {
	//创建数据库
	//连接数据库
	err := dao.InitMySQL()
	if err!=nil {
		panic(err)
	}
	//关闭数据库连接
	defer dao.Close()

	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	r:=routers.SetUpRouter()

	r.Run()
}
