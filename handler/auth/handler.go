package auth

import (
	"net/http"
	"telego/config"
	"telego/database"
	"telego/models"
	"telego/utils"

	"github.com/gin-gonic/gin"
)

var cfg *config.Config

func SetConfig(c *config.Config) {
	cfg = c
}

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
		return
	}
	if registerRequest.CheckPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Repeat Password is required"})
		return
	}

	if registerRequest.CheckPassword != registerRequest.Password {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid password check"})
		return
	}
	if registerRequest.DisplayName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "DisplayName is required"})
		return
	}
	if registerRequest.PhoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "PhoneNumber is required"})
		return
	}

	if database.DB.First(&models.User{}, "email = ?", registerRequest.Email).Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Email already exists"})
		return
	}
	if database.DB.First(&models.User{}, "username = ?", registerRequest.Username).Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Username already exists"})
		return
	}
	if database.DB.First(&models.User{}, "phone = ?", registerRequest.PhoneNumber).Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Phone number already exists"})
		return
	}

	var user models.User
	hashPassword, _ := utils.HashPassword(registerRequest.Password)
	user.Username = registerRequest.Username
	user.Email = registerRequest.Email
	user.Phone = registerRequest.PhoneNumber
	user.DisplayName = registerRequest.DisplayName
	user.AvatarURL = ""
	user.IsActive = true
	user.IsPrivate = false
	user.PasswordHash = hashPassword
	err := database.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Success", "data": registerRequest})
}

func Login(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request"})
		return
	}

	if loginRequest.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Username is required"})
		return
	}
	if loginRequest.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Password is required"})
		return
	}

	var user models.User
	err := database.DB.First(&user, "username = ?", loginRequest.Username).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "User not found"})
		return
	}

	if !utils.CheckPassword(loginRequest.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Invalid password"})
		return
	}

	token, err := utils.GenerateToken(user.Username, cfg.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login successful",
		"data": gin.H{
			"token": token,
			"user": gin.H{
				"id":           user.ID,
				"username":     user.Username,
				"display_name": user.DisplayName,
			},
		},
	})
}
