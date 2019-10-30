package server

import (
	user "github.com/empenguin1186/gin-gorm-tutorial/controller"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/users")
	{
		ctrl := user.Controller{}
		u.GET("", ctrl.Index)
		u.GET("/:id", ctrl.Show)
		u.POST("", ctrl.Create)
		u.PUT("", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	return r
}
