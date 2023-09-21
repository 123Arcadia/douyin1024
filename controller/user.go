package controller

import (
	"fmt"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/middlewares"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/models"
	response "github.com/123Arcadia/douyin1024CodeSpaceDemo.git/response"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

var (
	userService service.UserService
)

// url参数: user_id用户id;token用户鉴权token
func UserInfo(c *gin.Context) {
	userId := c.Query("user_id") //这里是string类型
	fmt.Println("user_id =", userId, ", ", c.Request.URL.Query())
	fmt.Println("keys= ", c.Keys)
	// 化为64位int64
	queryUserId, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		response.GetUserInfoFailed(c, "userId转换失败 /user/")
		return
	}
	// 获取suer信息
	user, getUserInfoErr := models.GetUserInfoByUserId(int64(queryUserId))
	if getUserInfoErr != nil {
		log.Println("查询用户信息错误!", getUserInfoErr)
		response.GetUserInfoFailed(c, "查询用户信息错误 /user/")
		return
	}
	response.GetUserInfoSuccess(c, 0, queryUserId, user)
}

func Register(c *gin.Context) {
	// 获取请求参数
	username := c.Query("username")
	password := c.Query("password")
	user, err := userService.Register(username, password)
	if err != nil {
		response.RegisterUserFailed(c, "用户注册失败")
		return
	}
	// 生产token
	token, err := middlewares.GenerateToken(user.ID, user.UserName, user.PassWord)
	if err != nil {
		log.Println("用户{%v}生成token异常", username)
		response.TokenGenernateFailed(c, "用户"+username+"生成token异常")
		return
	}
	// 注册成功
	response.RegisterUserSuccess(c, int32(user.ID), token)
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	// 登录
	user, err := userService.Login(username, password)
	if err != nil {
		response.UserLoginFailed(c, "用户登录失败")
		return
	}

	// 生成token
	token, err := middlewares.GenerateToken(user.ID, username, password)
	if err != nil {
		response.UserLoginFailed(c, "token生成失败")
		return
	}
	response.UserLoginSuccess(c, int32(user.ID), token)
}
