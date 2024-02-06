package controller

import (
	"fmt"
	"net/http"

	"github.com/funnyzak/gogin/model"
	"github.com/gin-gonic/gin"
)

func PageNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound,
		&model.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("Resource not found: %s", c.Request.RequestURI),
		})
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{})
}
