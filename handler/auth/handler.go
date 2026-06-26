package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerRequest RegisterRequest
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request"})
		return
	}

	if registerRequest.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Email is required"})
		return
	}
	if registerRequest.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Username is required"})
		return
	}
	if registerRequest.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Password is required"})
	}
	if registerRequest.CheckPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Repeat Password is required"})
		return
	}

	if registerRequest.CheckPassword != registerRequest.Password {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid password check"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Success", "data": registerRequest})
}
