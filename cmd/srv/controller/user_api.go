package controller

import (
	"github.com/gin-gonic/gin"
)

type userAPI struct {
	r gin.IRouter
}

func (ua *userAPI) serve() {
	ur := ua.r.Group("")
	ur.POST("/login", ua.login)
	ur.POST("/register", ua.register)

	v1 := ua.r.Group("v1")
	{
		apiv1 := &apiV1{r: v1}
		apiv1.serve()
	}
}

func (ua *userAPI) login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login",
	})
}

func (ua *userAPI) register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register",
	})
}
