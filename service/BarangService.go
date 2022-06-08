package service

import (
	"log"
	"net/http"
	"restapi-tugas/modals"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Get Barang
func GetAllBarang(c *gin.Context) {
	Barangs := modals.FindAllBarangs()
	log.Println("Barangs -> ", Barangs)
	c.JSON(http.StatusOK, Barangs)
}

// Get One Barang
func GetOneBarang(c *gin.Context) {
	Barang := modals.FindByBarangId(c.Param("id"))
	if Barang.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	log.Println("Barang -> ", Barang)
	c.JSON(http.StatusOK, Barang)
}

// Post Barang

func PostBarang(c *gin.Context) {
	barang := modals.Barang{}
	session := sessions.Default(c)
	sessionID := session.Get("session_id")
	if sessionID == nil {
		log.Fatalf("error")
	}

	barang.IdUser = sessionID.(uint64)
	timenow := time.Now()
	barang.CreatedAt = timenow
	barang.UpdatedAt = timenow
	err := c.BindJSON(&barang)
	if err != nil {
		c.String(400, ":%s", err.Error())
		return
	}
	newBarang := modals.CreateBarang(barang)
	c.JSON(http.StatusCreated, newBarang)
}

// delete Barang

func DeleteBarang(c *gin.Context) {
	isDelete := modals.DeleteBarang(c.Param("id"))
	if isDelete {
		c.JSON(http.StatusOK, "Successfully")
		return
	}
	c.JSON(http.StatusNotFound, "Error")
}
func PutBarang(c *gin.Context) {
	Barang := modals.Barang{}
	//GET ID
	session := sessions.Default(c)
	sessionID := session.Get("session_id")
	if sessionID == nil {
		log.Fatalf("error")
	}

	Barang.IdUser = sessionID.(uint64)
	timenow := time.Now()
	Barang.UpdatedAt = timenow
	//BIND JSON
	err := c.BindJSON(&Barang)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
		return
	}

	Barang = modals.UpdateBarang(c.Param("id"), Barang)
	if Barang.Id == 0 {
		c.JSON(http.StatusNotFound, "Errorbro")
		return
	}
	c.JSON(http.StatusOK, "Successfully")
}
