package controllers

import (
	"ecomresto/initializers"
	"ecomresto/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var body struct {
		Username  string `form:"username" binding:"required"`
		Email     string `form:"email" binding:"required,email"`
		Password  string `form:"password" binding:"required"`
		CPassword string `form:"c_password" binding:"required"`
		Phone     int    `form:"phone" binding:"required"`
	}

	if err := c.ShouldBind(&body); err != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"error": "Failed to read body"})
		return
	}

	user := models.Users{
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password, // Hash the password before saving
		Phone:    body.Phone,
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"error": "Failed to create user"})
		return
	}

	c.HTML(http.StatusOK, "signup.html", gin.H{"message": "User created successfully"})
}
