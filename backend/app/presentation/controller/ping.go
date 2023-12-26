package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct{}

func (*PingController) GetPing(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
		"message": "ping",
	})
}