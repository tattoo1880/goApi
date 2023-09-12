package router

import (
	"fmt"
	"tattoo_code/app/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		fmt.Println(err)
	}
	// api路由组
	api := r.Group("/api")
	{
		api.GET("/ping", controller.Ping)
		api.POST("/info", controller.CreateInfo)
		api.GET("/info/:id", controller.GetInfoById)
		api.PUT("/info/:id", controller.UpdateInfo)
		api.DELETE("/info/:id", controller.DeleteInfo)
	}
	return r
}
