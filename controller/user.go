package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	g "main/global"
	"main/modal"
	"main/services"
	"main/utils"
)

const (
	HasExist     = "user has existed"
	GenerateFail = "generate user failed"
	WrongPsw     = "wrong password"
)

func Enroll(c *gin.Context) {
	//1.接受参数并校验
	p := new(modal.ParamUser)
	if err := c.ShouldBind(p); err != nil {
		utils.RespOK(c, &utils.Resp{
			Msg:    "wrong parameter",
			Code:   1001,
			Status: "failed",
		})
		return
	}
	if p.Username == "" || p.Password == "" || p.RePassword != p.Password {
		utils.RespOK(c, &utils.Resp{
			Msg:    "Required content is empty or the two passwords are inconsistent",
			Code:   1001,
			Status: "failed",
		})
		return
	}
	//2.用户注册
	err := services.Enroll(p)
	if err != nil {
		if fmt.Sprintf("%v", err) == HasExist {
			utils.RespOK(c, &utils.Resp{
				Msg:    "user already exists",
				Code:   1001,
				Status: "failed",
			})
			return
		}
		if fmt.Sprintf("%v", err) == GenerateFail {
			utils.RespFailed(c, &utils.Resp{
				Msg:    "generate user failed",
				Code:   1002,
				Status: "failed",
			})
			g.Logger.Warn("generate user failed")
			return
		}
		utils.RespFailed(c, &utils.Resp{
			Msg:    "internal error",
			Code:   1002,
			Status: "failed",
		})
		g.Logger.Warn(fmt.Sprintf("err:%v", err))
		return
	}
	//3.返回响应
	utils.RespOK(c, &utils.Resp{
		Msg:    "enroll successfully",
		Code:   1000,
		Status: "successful",
	})
	g.Logger.Info(fmt.Sprintf("user:%v enroll ", p.Username))
}

func Login(c *gin.Context) {
	//1.参数获取与校验
	p := new(modal.ParamUser)
	if err := c.ShouldBind(p); err != nil {
		utils.RespOK(c, &utils.Resp{
			Msg:    "wrong parameter",
			Code:   1001,
			Status: "failed",
		})
	}
	if p.Username == "" || p.Password == "" {
		utils.RespOK(c, &utils.Resp{
			Msg:    "Required content is empty",
			Code:   1001,
			Status: "failed",
		})
	}
	//2.用户登录
	tokenStr, err := services.Login(p)
	if err != nil {
		if fmt.Sprintf("%v", err) == WrongPsw {
			utils.RespOK(c, &utils.Resp{
				Msg:    "wrong password",
				Code:   1001,
				Status: "failed",
			})
			return
		}
		utils.RespFailed(c, &utils.Resp{
			Msg:    "internal error",
			Code:   1003,
			Status: "failed",
		})
		return
	}
	//3.返回响应
	utils.RespOK(c, &utils.Resp{
		Msg:    "login successfully",
		Code:   1000,
		Status: "successfully",
		Data:   tokenStr,
	})
	g.Logger.Info(fmt.Sprintf("user:%v login", p.Username))
}

func RevisePassword(c *gin.Context) {
	//获用户信息
	uid, ok := utils.GetCurrentUser(c)
	if !ok {
		utils.RespOK(c, &utils.Resp{
			Msg:  "not login",
			Code: 1001,
		})
		return
	}
	//获取参数并校验
	p := new(modal.ParamUser)
	if err := c.ShouldBind(p); err != nil {
		utils.RespOK(c, &utils.Resp{
			Msg:    "wrong parameter",
			Code:   1001,
			Status: "failed",
		})
		return
	}
	if p.Password == "" || p.OriPassword == "" || p.RePassword != p.Password {
		utils.RespOK(c, &utils.Resp{
			Msg:    "Required content is empty or the two passwords are inconsistent",
			Code:   1001,
			Status: "failed",
		})
		return
	}
	if p.OriPassword == p.Password {
		utils.RespOK(c, &utils.Resp{
			Msg:    "The new password cannot be the same as the original password",
			Code:   1001,
			Status: "failed",
		})
		return
	}
	//修改密码
	err := services.RevisePwd(p, uid)
	if err != nil {
		if fmt.Sprintf("%v", err) == WrongPsw {
			utils.RespOK(c, &utils.Resp{
				Msg:    "wrong password",
				Code:   1001,
				Status: "failed",
			})
			return
		}
		utils.RespFailed(c, &utils.Resp{
			Msg:    "internal error",
			Code:   1003,
			Status: "failed",
		})
		g.Logger.Warn(fmt.Sprintf("%v", err))
		return
	}
	//返回响应
	utils.RespOK(c, &utils.Resp{
		Msg:    "revise password successfully",
		Code:   1000,
		Status: "successfully",
	})
	g.Logger.Info(fmt.Sprintf("userID:%v revise his password", uid))
}
