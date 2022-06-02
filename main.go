package main

import (
	"io"
	"os"
	"restapi-tugas/database"
	"restapi-tugas/middlewares"
	"restapi-tugas/modals"
	"restapi-tugas/src"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func setupLoggerOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLoggerOutput() // setup logging

	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("userpasd", middlewares.UserPasd)

	}

	router.Use(gin.Recovery(), middlewares.Logger()) // logging

	v1 := router.Group("/v1")
	src.AddUserRouter(v1)
	src.AddBarangRouter(v1)
	src.AddGudangRouter(v1)

	go func() {
		// 加入 MySQL Connection
		database.DD()
		database.DBconnect.AutoMigrate(&modals.User{}, &modals.Barang{}, &modals.Gudang{})

	}()

	router.Run(":8000")
}
