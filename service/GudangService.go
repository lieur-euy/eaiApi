package service

import (
	"log"
	"net/http"
	"restapi-tugas/modals"

	"github.com/gin-gonic/gin"
)

// Get Gudang
func GetAllGudang(c *gin.Context) {
	Gudangs := modals.FindAllGudangs()
	log.Println("Gudangs -> ", Gudangs)
	c.JSON(http.StatusOK, Gudangs)
}

// Get One Gudang
func GetOneGudang(c *gin.Context) {
	Gudang := modals.FindByGudangId(c.Param("id"))
	if Gudang.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	log.Println("Gudang -> ", Gudang)
	c.JSON(http.StatusOK, Gudang)
}

// Post Gudang

func PostGudang(c *gin.Context) {
	Gudang := modals.Gudang{}
	err := c.BindJSON(&Gudang)
	if err != nil {
		c.String(400, ":%s", err.Error())
		return
	}
	newGudang := modals.CreateGudang(Gudang)
	c.JSON(http.StatusCreated, newGudang)
}

// delete Gudang

func DeleteGudang(c *gin.Context) {
	isDelete := modals.DeleteGudang(c.Param("id"))
	if isDelete {
		c.JSON(http.StatusOK, "Successfully")
		return
	}
	c.JSON(http.StatusNotFound, "Error")
}

// put Gudang

func PutGudang(c *gin.Context) {
	Gudang := modals.Gudang{}
	err := c.BindJSON(&Gudang)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
		return
	}
	Gudang = modals.UpdateGudang(c.Param("id"), Gudang)
	if Gudang.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, "Successfully")
}
