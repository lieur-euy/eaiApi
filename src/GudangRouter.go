package src

import (
	session "restapi-tugas/middlewares"
	"restapi-tugas/service"

	"github.com/gin-gonic/gin"
)

func AddGudangRouter(r *gin.RouterGroup) {
	Gudang := r.Group("/Gudangs", session.SetSession())
	Gudang.GET("/", service.GetAllGudang)
	Gudang.GET("/:id", service.GetOneGudang)
	Gudang.POST("/register", service.PostGudang)
	Gudang.PUT("/update/:id", service.PutGudang)
}
