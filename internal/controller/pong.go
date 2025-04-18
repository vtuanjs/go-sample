package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct {
}

func NewPongController() *PongController {
	return &PongController{}
}

func (pc *PongController) Pong(c *gin.Context) {
	name := c.Param("name")
	// c.ShouldBindJSON()
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
		"users":   []string{"user1", "user2"},
	})
}
