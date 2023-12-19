package dao

import "github.com/jinzhu/gorm"

var(
	DB *gorm.DB
)

func InitMySQL()(err error){
	dsn := "root:hyx20040731@tcp(localhost:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB,err = gorm.Open("mysql",dsn)//注意此处不要写:=，否则成了在这个函数内部新声明了一个变量DB
	if err!=nil {
		return
	}
	return DB.DB().Ping()//如果ping的通，则返回一个nil
}

func Close(){
	DB.Close()
}