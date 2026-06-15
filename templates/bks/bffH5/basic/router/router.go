package router

import (
	"bks/bffH5/basic/middlewares"
	"bks/bffH5/handler/api"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.POST("/upload", api.Upload)
	r.POST("/pay/notify", api.Notify)
	group := r.Group("/v1")
	group.Use(middlewares.MiddleWares())

	group.POST("/users/add", api.UserAdd)
	group.POST("/users/login", api.UserLogin)
	group.POST("/users/info", api.UserInfo)
	group.POST("/users/list", api.UserList)
	group.POST("/users/update", api.UserUpdate)
	group.POST("/users/delete", api.UserDelete)
}
