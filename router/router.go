package router

import (
	"github.com/gin-gonic/gin"
	"picking/handler"
)

func InitUserRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/api/get_user", handler.GetUser)
	router.POST("/api/add_user", handler.AddUser)
	router.DELETE("/api/del_user", handler.DelUser)
	router.GET("/api/get_user_list", handler.GetUserList)
	return router
}
