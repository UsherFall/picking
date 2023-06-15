package router

import (
	"github.com/gin-gonic/gin"
	"picking/handler"
	"picking/middleware"
)

func InitUserRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.GET("/api/get_user", handler.GetUser)
	router.POST("/api/add_user", handler.AddUser)
	router.DELETE("/api/del_user", handler.DelUser)
	router.GET("/api/get_user_list", handler.GetUserList)
	return router
}
