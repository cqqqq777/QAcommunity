package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	g "main/global"
	"main/modal"
	"main/services"
	"main/utils"
	"strconv"
)

func GetAllPostsTitle(c *gin.Context) {
	title, err := services.GetAllPostTitle()
	if err != nil {
		utils.RespFailed(c, &utils.Resp{
			Msg:    "internal error",
			Code:   1003,
			Status: "failed",
		})
		g.Logger.Warn(fmt.Sprintf("an error occurred while querying all post title  error:%v", err))
		return
	}
	utils.RespOK(c, &utils.Resp{
		Code:   1000,
		Status: "successfully",
		Data:   title,
	})
}

func PostDetail(c *gin.Context) {
	//获取帖子id并校验
	postID := c.PostForm("post-id")
	id, err := strconv.ParseInt(postID, 10, 64)
	if postID == "" || err != nil {
		utils.RespOK(c, &utils.Resp{
			Msg:    "wrong format",
			Code:   1001,
			Status: "failed",
		})
		return
	}
	//查询帖子细节
	post, err := services.PostDetail(int(id))
	if err != nil {
		utils.RespFailed(c, &utils.Resp{
			Msg:    "internal error",
			Code:   1003,
			Status: "failed",
		})
		g.Logger.Warn(fmt.Sprintf("an error occurred while querying the post   error:%v", err))
		return
	}
	//返回响应
	data := fmt.Sprintf("{\"title\":\"%v\",\"content\":\"%v\",\"createAt\":\"%v\",\"updateAt\":\"%v\"}",
		post.Title,
		post.Content,
		utils.GetFormatTime(post.CreateAt),
		utils.GetFormatTime(post.UpdateAt),
	)
	utils.RespOK(c, &utils.Resp{
		Code:   1000,
		Data:   data,
		Status: "successfully",
	})
}

func QueryPostComment(c *gin.Context) {
	//获取帖子id
	ID := c.Query("post-id")
	postID, err := strconv.ParseInt(ID, 10, 64)
	if ID == "" || err != nil {
		utils.RespOK(c, &utils.Resp{
			Msg:    "wrong format",
			Code:   1001,
			Status: "failed",
		})
		return
	}
	//查询帖子下的回复
	comments, err := services.QueryPostComment(int(postID))
	if err != nil {
		utils.RespFailed(c, &utils.Resp{
			Msg:    "internal error",
			Code:   1003,
			Status: "failed",
		})
		g.Logger.Warn(fmt.Sprintf("an error occurred while querying the post'comments   error:%v", err))
		return
	}
	//返回响应
	utils.RespFailed(c, &utils.Resp{
		Data:   comments,
		Code:   1000,
		Status: "successfully",
	})
}

func PublishPost(c *gin.Context) {
	//获取发帖人的id
	authorID, ok := utils.GetCurrentUser(c)
	if !ok {
		utils.RespOK(c, &utils.Resp{
			Msg:    "need login",
			Code:   1001,
			Status: "failed",
		})
		return
	}
	//获取帖子参数并校验
	postParam := new(modal.ParamPost)
	if err := c.ShouldBind(postParam); err != nil {
		utils.RespOK(c, &utils.Resp{
			Msg:    "wrong parameter",
			Code:   1001,
			Status: "failed",
		})
		return
	}
	if err := services.PublishPost(postParam, authorID); err != nil {
		utils.RespFailed(c, &utils.Resp{
			Msg:    "internal error",
			Code:   1003,
			Status: "failed",
		})
		g.Logger.Warn(fmt.Sprintf("an error occurred while publish the post   error:%v", err))
		return
	}
	utils.RespOK(c, &utils.Resp{
		Code:   1000,
		Status: "successfully",
	})
	g.Logger.Info(fmt.Sprintf("userid:%v publish a post", authorID))
}

func GetPersonalPost(c *gin.Context) {
	//获取用户id
	authorID, ok := utils.GetCurrentUser(c)
	if !ok {
		utils.RespOK(c, &utils.Resp{
			Msg:  "need login",
			Code: 1001,
		})
		return
	}
	posts, err := services.GetPersonalPost(authorID)
	if err != nil {
		utils.RespFailed(c, &utils.Resp{
			Msg:    "internal error",
			Code:   1003,
			Status: "failed",
		})
		g.Logger.Warn(fmt.Sprintf("an error occurred while query the personal post  error:%v", err))
		return
	}
	utils.RespOK(c, &utils.Resp{
		Data:   posts,
		Code:   1000,
		Status: "successfully",
	})
}

func UpdatePost(c *gin.Context) {
	//获取发帖人的id
	authorID, ok := utils.GetCurrentUser(c)
	if !ok {
		utils.RespOK(c, &utils.Resp{
			Msg:    "need login",
			Code:   1001,
			Status: "failed",
		})
		return
	}
	//获取帖子参数并校验
	postParam := new(modal.ParamPost)
	if err := c.ShouldBind(postParam); err != nil {
		utils.RespOK(c, &utils.Resp{
			Msg:    "wrong parameter",
			Code:   1001,
			Status: "failed",
		})
		return
	}
	//处理更新请求
	if err := services.UpdatePost(postParam, authorID); err != nil {
		utils.RespFailed(c, &utils.Resp{
			Msg:    "internal error",
			Code:   1003,
			Status: "failed",
		})
		g.Logger.Warn(fmt.Sprintf("an error occurred while update post  error:%v", err))
		return
	}
	//返回响应
	utils.RespOK(c, &utils.Resp{
		Code:   1000,
		Status: "successfully",
	})
	g.Logger.Info(fmt.Sprintf("userid:%v update postid:%v", authorID, postParam.PostID))
}

func DeletePost(c *gin.Context) {
	//获取删除帖子的人的id
	authorID, ok := utils.GetCurrentUser(c)
	if !ok {
		utils.RespOK(c, &utils.Resp{
			Msg:    "need login",
			Code:   1001,
			Status: "failed",
		})
		return
	}
	ID := c.PostForm("post-id")
	postID, err := strconv.ParseInt(ID, 10, 64)
	if ID == "" || err != nil {
		utils.RespOK(c, &utils.Resp{
			Msg:    "wrong parameter",
			Code:   1001,
			Status: "failed",
		})
		return
	}
	if err := services.DeletePost(authorID, int(postID)); err != nil {
		utils.RespFailed(c, &utils.Resp{
			Msg:    "internal error",
			Code:   1003,
			Status: "failed",
		})
		g.Logger.Warn(fmt.Sprintf("an error occurred while update post  error:%v", err))
		return
	}
	utils.RespOK(c, &utils.Resp{
		Code:   1000,
		Status: "successfully",
	})
	g.Logger.Info(fmt.Sprintf("userID:%v delete postID:%v", authorID, postID))
}
