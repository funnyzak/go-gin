package controller

import (
	"fmt"
	"go-gin/internal/gogin"
	"go-gin/mappers"
	"go-gin/model"
	"go-gin/pkg/mygin"
	"go-gin/pkg/utils"
	"go-gin/service/singleton"
	"net/http"

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
	var loginForm mappers.LoginForm

	isPage := utils.ParseBool(c.Query("page"), false)
	showError := func(err error) {
		gogin.ShowErrorPage(c, mygin.ErrInfo{
			Title: "Login failed",
			Code:  http.StatusNotAcceptable,
			Msg:   err.Error(),
			Link:  fmt.Sprintf("%s/login", singleton.Conf.Site.BaseURL),
			Btn:   "Back to login",
		}, isPage)
	}

	if err := mygin.BindForm(c, isPage, &loginForm); err != nil {
		showError(err)
		return
	}

	u := new(model.User)
	token, err := u.Login(loginForm, singleton.DB, singleton.Conf)
	if err != nil {
		showError(err)
	} else {
		if isPage {
			gogin.UserLoginSuccess(c, token)
			c.Redirect(http.StatusFound, fmt.Sprintf("%s/user/mine", singleton.Conf.Site.BaseURL))
		} else {
			mygin.ResponseJSON(c, http.StatusOK, gin.H{
				"token": token,
				"user":  u,
			}, "Login success")
		}
	}
}

func (ua *userAPI) register(c *gin.Context) {
	var registerForm mappers.RegisterForm

	isPage := utils.ParseBool(c.Query("page"), false)
	showError := func(err error) {
		gogin.ShowErrorPage(c, mygin.ErrInfo{
			Title: "Register failed",
			Code:  http.StatusNotAcceptable,
			Msg:   err.Error(),
			Link:  fmt.Sprintf("%s/register", singleton.Conf.Site.BaseURL),
			Btn:   "Back to home",
		}, isPage)
	}

	if err := mygin.BindForm(c, isPage, &registerForm); err != nil {
		showError(err)
		return
	}

	u := new(model.User)
	err := u.Register(registerForm, singleton.DB, singleton.Conf)
	if err != nil {
		showError(err)
	} else {
		if isPage {
			c.Redirect(http.StatusFound, fmt.Sprintf("%s/login", singleton.Conf.Site.BaseURL))
		} else {
			mygin.ResponseJSON(c, http.StatusOK, u, "Register success")
		}
	}
}
