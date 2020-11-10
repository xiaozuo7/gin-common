package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type Response struct {
	Code string `json:"code"`
	Data gin.H  `json:"data"`
	Msg  string `json:"msg"`
}

type RenderFunc func(c *gin.Context)

type JSONRenderer interface {
	RenderJSONFunc() RenderFunc
}

func (e *Response) RenderJSONFunc() RenderFunc {
	return func(c *gin.Context) {
		c.Render(http.StatusOK, render.JSON{Data: e})
	}
}
