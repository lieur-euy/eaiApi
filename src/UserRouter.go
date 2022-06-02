package src

import (
	session "restapi-tugas/middlewares"
	"restapi-tugas/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users", session.SetSession())
	user.GET("/", service.GetAllUser)
	user.GET("/:id", service.GetOneUser)
	user.POST("/register", service.PostUser)
	user.POST("/login", service.LoginUser)
	user.PUT("/update/:id", service.PutUser)

	user.Use(session.AuthSession())
	{
		user.DELETE("/:id", service.DeleteUser)
		user.GET("/session", service.GetSession)
		user.GET("/logout", service.LogoutUser)
	}
}
