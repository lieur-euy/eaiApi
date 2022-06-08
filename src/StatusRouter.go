package src

import (
	session "restapi-tugas/middlewares"
	"restapi-tugas/service"

	"github.com/gin-gonic/gin"
)

func AddStatusRouter(r *gin.RouterGroup) {
	Status := r.Group("/Status", session.SetSession())
	Status.GET("/", service.GetAllStatus)
	Status.GET("/:id", service.GetOneStatus)
	Status.POST("/input", service.PostStatus)
	Status.PUT("/update/:id", service.PutStatus)
	Status.DELETE("delete/:id", service.DeleteStatus)

}
