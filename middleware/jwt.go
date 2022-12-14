package middleware

import (
	"github.com/gin-gonic/gin"
	"main/modal"
	"main/utils"
	"strings"
)

func JWTAuth(c *gin.Context) {
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	//假设Token放在Header的Authorization中，并使用Bearer开头
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		utils.RespOK(c, &utils.Resp{
			Msg:    "empty header",
			Code:   1001,
			Status: "failed",
		})
		c.Abort()
		return
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		utils.RespOK(c, &utils.Resp{
			Msg:    "wrong format",
			Code:   1001,
			Status: "failed",
		})
		c.Abort()
		return
	}
	myClaim, err := utils.ParseToken(parts[1])
	if err != nil {
		utils.RespOK(c, &utils.Resp{
			Msg:    "invalid token",
			Code:   1001,
			Status: "failed",
		})
	}
	c.Set(modal.CtxGetUID, myClaim.UserID)
}
