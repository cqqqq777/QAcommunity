package boot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/controller"
	g "main/global"
	"main/middleware"
	"time"
)

func InitRouters() {
	r := gin.New()
	r.Use(middleware.GinLogger(g.Logger), middleware.GinRecovery(g.Logger, true))
	r.Use(middleware.Cors)
	PublicGroup := r.Group("/api/public")
	{
		PublicGroup.POST("/enroll", controller.Enroll)
		PublicGroup.POST("/login", controller.Login)
		PublicGroup.GET("/postTitle", controller.GetAllPostsTitle)
		PublicGroup.POST("/post", controller.PostDetail)
		PublicGroup.GET("/postComments", controller.QueryPostComment)
	}
	PrivateGroup := r.Group("/api/private", middleware.JWTAuth)
	{
		PrivateGroup.PUT("/revisePwd", controller.RevisePassword)
		PrivateGroup.POST("/post", controller.PublishPost)
		PrivateGroup.GET("/post", controller.GetPersonalPost)
		PrivateGroup.PUT("/post", controller.UpdatePost)
		PrivateGroup.DELETE("/post", controller.DeletePost)
		PrivateGroup.GET("/comment")
		PrivateGroup.POST("/comment")
		PrivateGroup.DELETE("/comment")
		PrivateGroup.POST("/likePost")
		PrivateGroup.POST("/likeComment")
		PrivateGroup.DELETE("/likePost")
		PrivateGroup.DELETE("/likeComment")
		PrivateGroup.POST("/follow")
		PrivateGroup.DELETE("/follow")
	}
	if err := r.Run(); err != nil {
		g.Logger.Fatal(fmt.Sprintf("boot service failed err:%v", err))
	}
	g.Logger.Info(fmt.Sprintf("sevice boot successfully at %v", time.Now().Format("2006/05/04 15:02:01")))
}
