package utils

import (
	"github.com/gin-gonic/gin"
	"main/modal"
)

func GetCurrentUser(c *gin.Context) (UserID int, ok bool) {
	uid, ok := c.Get(modal.CtxGetUID)
	if !ok {
		return
	}
	UserID, ok = uid.(int)
	if !ok {
		return
	}
	return
}
