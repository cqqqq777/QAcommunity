package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	Msg    interface{} `json:"msg,omitempty"`
	Code   int         `json:"code,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Status string      `json:"status,omitempty"`
}

func RespOK(c *gin.Context, r *Resp) {
	c.JSON(http.StatusOK, r)
}

func RespFailed(c *gin.Context, r *Resp) {
	c.JSON(http.StatusInternalServerError, r)
}
