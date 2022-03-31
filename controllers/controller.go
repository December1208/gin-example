package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct{}

func (ctrl Controller) Index(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    "123",
		"success": true,
		"detail":  "",
	})
}


