package server

import (
	user "github.com/empenguin1186/gin-gorm-tutorial/controller"
	"github.com/empenguin1186/gin-gorm-tutorial/repository"
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
		repository := repository.NewUserRepositoryOnMemory()
		service := user.NewService(repository)
		ctrl := user.NewController(service)

		u.GET("", ctrl.Index)
		u.GET("/:id", ctrl.Show)
		u.POST("", ctrl.Create)
		u.PUT("", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	return r
}
