package mygin

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int         `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

type ErrInfo struct {
	Code  int    `json:"code,omitempty"`
	Title string `json:"title,omitempty"`
	Msg   string `json:"msg,omitempty"`
	Link  string `json:"link,omitempty"`
	Btn   string `json:"btn,omitempty"`
}

func ResponseJSON(c *gin.Context, code int, data interface{}, messages ...string) {
	rlt := &Response{}
	if code > 0 {
		rlt.Code = code
	}
	if data != nil {
		rlt.Data = data
	}
	if len(messages) > 0 {
		rlt.Message = messages[0]
	}
	c.AbortWithStatusJSON(code, rlt)
}
