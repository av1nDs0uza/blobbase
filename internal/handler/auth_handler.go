package handler

import (
	"errors"
	"net/http"

	appErr "filestorage/internal/errors"
	"filestorage/internal/service"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	err := service.RegisterUser(body.Email, body.Password)
	if err != nil {

		if errors.Is(err, appErr.ErrEmailExists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created"})
}
