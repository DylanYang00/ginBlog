package main

import (
	"ginBlog/model"
	"ginBlog/routers"
)

func main() {
	//初始化数据库
	model.InitDb()

	routers.InitRouter()

}
