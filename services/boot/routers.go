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
	PublicGroup := r.Group("/public")
	{
		PublicGroup.POST("/enroll", controller.Enroll)
		PublicGroup.POST("/login", controller.Login)
		PublicGroup.GET("/problem")
	}
	PrivateGroup := r.Group("/private", middleware.JWTAuth)
	{
		PrivateGroup.PUT("/revisePwd", controller.RevisePassword)
		PrivateGroup.POST("/problem")
		PrivateGroup.GET("/problem")
		PrivateGroup.PUT("/problem")
		PrivateGroup.DELETE("/problem")
		PrivateGroup.GET("/comment")
		PrivateGroup.PUT("/comment")
		PrivateGroup.POST("/comment")
		PrivateGroup.DELETE("/comment")
		PrivateGroup.POST("/likeProblem")
		PrivateGroup.POST("/likeComment")
		PrivateGroup.DELETE("/likeProblem")
		PrivateGroup.DELETE("/likeComment")
		PrivateGroup.POST("/follow")
		PrivateGroup.DELETE("/follow")
	}
	if err := r.Run(); err != nil {
		g.Logger.Fatal(fmt.Sprintf("boot service failed err:%v", err))
	}
	g.Logger.Info(fmt.Sprintf("sevice boot successfully at %v", time.Now().Format("2006/05/04 15:02:01")))
}
