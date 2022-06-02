package src

import (
	session "restapi-tugas/middlewares"
	"restapi-tugas/service"

	"github.com/gin-gonic/gin"
)

func AddBarangRouter(r *gin.RouterGroup) {
	Barang := r.Group("/Barang", session.SetSession())
	Barang.GET("/", service.GetAllBarang)
	Barang.GET("/:id", service.GetOneBarang)
	Barang.POST("/input", service.PostBarang)
	Barang.PUT("/update/:id", service.PutBarang)
	Barang.DELETE("delete/:id", service.DeleteBarang)

}
