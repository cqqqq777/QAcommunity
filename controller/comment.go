package controller

import (
	"github.com/gin-gonic/gin"
	"main/modal"
	"main/services"
	"main/utils"
)

func CommentPost(c *gin.Context) {
	comment := new(modal.ParamComment)
	if err := c.ShouldBind(comment); err != nil {
		utils.RespOK(c, &utils.Resp{
			Msg:  "wrong format",
			Code: 1001,
		})
		return
	}
	UserID, ok := c.Get(modal.CtxGetUID)
	if !ok {
		utils.RespOK(c, &utils.Resp{
			Msg:  "need login",
			Code: 1001,
		})
		return
	}
	authorID, ok := UserID.(int)
	if !ok {
		utils.RespOK(c, &utils.Resp{
			Msg:  "need login",
			Code: 1001,
		})
		return
	}
	comment.AuthorID = authorID
	err := services.CommentPost(comment)
}
