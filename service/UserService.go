package service

import (
	session "restapi-tugas/middlewares"
	"restapi-tugas/modals"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get User
func GetAllUser(c *gin.Context) {
	users := modals.FindAllUsers()
	log.Println("Users -> ", users)
	c.JSON(http.StatusOK, users)
}

// Get One User
func GetOneUser(c *gin.Context) {
	user := modals.FindByUserId(c.Param("id"))
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	log.Println("User -> ", user)
	c.JSON(http.StatusOK, user)
}

// Post User

func PostUser(c *gin.Context) {
	user := modals.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.String(400, ":%s", err.Error())
		return
	}
	newUser := modals.CreateUser(user)
	c.JSON(http.StatusCreated, newUser)
}

// delete User

func DeleteUser(c *gin.Context) {
	isDelete := modals.DeleteUser(c.Param("id"))
	if isDelete {
		c.JSON(http.StatusOK, "Successfully")
		return
	}
	c.JSON(http.StatusNotFound, "Error")
}

// put User

func PutUser(c *gin.Context) {
	user := modals.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
		return
	}
	user = modals.UpdateUser(c.Param("id"), user)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, "Successfully")
}

// Create Users

// LoginUser
func LoginUser(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	user := modals.LoginUser(name, password)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error bro")
		return
	}
	session.SaveSession(c, user.Id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged in",
		"user":    user,
		"Session": session.GetSession(c),
	})
}

// User get Session
func GetSession(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": session.GetSession(c),
	})
}

// Logout User
func LogoutUser(c *gin.Context) {
	session.ClearSession(c)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}
