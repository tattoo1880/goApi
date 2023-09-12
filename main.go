package main

import (
	"tattoo_code/app/models"
	"tattoo_code/database"
	"tattoo_code/router"
)

func main() {
	// automigrate数据库
	err := database.DB.AutoMigrate(&models.Info{})
	if err != nil {
		panic(err)
	}
	// 初始化路由
	r := router.InitRouter()
	// port := os.Getenv("HTTP_PORT")
	r.Run()
}
