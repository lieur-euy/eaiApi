package service

import (
	"log"
	"net/http"
	"restapi-tugas/modals"

	"github.com/gin-gonic/gin"
)

// Get Status
func GetAllStatus(c *gin.Context) {
	Statuss := modals.FindAllStatuss()
	log.Println("Statuss -> ", Statuss)
	c.JSON(http.StatusOK, Statuss)
}

// Get One Status
func GetOneStatus(c *gin.Context) {
	Status := modals.FindByStatusId(c.Param("id"))
	if Status.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	log.Println("Status -> ", Status)
	c.JSON(http.StatusOK, Status)
}

// Post Status

func PostStatus(c *gin.Context) {
	Status := modals.Status{}
	err := c.BindJSON(&Status)
	if err != nil {
		c.String(400, ":%s", err.Error())
		return
	}
	newStatus := modals.CreateStatus(Status)
	c.JSON(http.StatusCreated, newStatus)
}

// delete Status

func DeleteStatus(c *gin.Context) {
	isDelete := modals.DeleteStatus(c.Param("id"))
	if isDelete {
		c.JSON(http.StatusOK, "Successfully")
		return
	}
	c.JSON(http.StatusNotFound, "Error")
}

// put Status

func PutStatus(c *gin.Context) {
	Status := modals.Status{}
	err := c.BindJSON(&Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
		return
	}
	Status = modals.UpdateStatus(c.Param("id"), Status)
	if Status.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, "Successfully")
}
