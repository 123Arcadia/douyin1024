package router

import (
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/controller"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/public", "./public")
	r.StaticFile("/favicon.ico", "./public/favicon.ico") // StaticFile 方法用于提供一个特定文件。
	r.LoadHTMLGlob("templates/*")

	// home page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "抖音极简版",
		})
	})

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/test", controller.Test)
	apiRouter.GET("/feed", controller.Feed)
	apiRouter.GET("/user/", middlewares.Auth(), controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", middlewares.UserPublishAuth(), controller.Publish)
	apiRouter.GET("/publish/list", middlewares.Auth(), controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", middlewares.Auth(), controller.FavoriteAction)
	apiRouter.GET("/favorite/list", middlewares.Auth(), controller.FavoriteList)
	apiRouter.POST("/comment/action/", middlewares.Auth(), controller.CommentAction)
	apiRouter.GET("/comment/list/", middlewares.Auth(), controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", middlewares.Auth(), controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", middlewares.Auth(), controller.FollowList)
	apiRouter.GET("/relation/follower/list/", middlewares.Auth(), controller.FollowerList)
	// 好友: 只要有关注关系的user(单方面也可以)
	apiRouter.GET("/relation/friend/list/", middlewares.Auth(), controller.FriendList)
	// 发送消息
	apiRouter.POST("/message/action/", middlewares.Auth(), controller.MessageAction)
	// 聊天记录
	apiRouter.GET("/message/chat/", middlewares.Auth(), controller.MessageChat)
}
