package controller

import (
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/models"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/response"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// url参数: user_id用户id;token用户鉴权token
func UserInfo(c *gin.Context) {
	//userId := c.Query("userId")这里是string类型
	userId := c.GetUint("userId")
	// 化为64位int64
	query_user_id, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		response.GetUserInfoFailed(c, "userId转换失败 /user/")
		return
	}
	// 获取suer信息
	user, getUserInfoErr := models.GetUserInfo(c, int64(query_user_id))
	if getUserInfoErr != nil {
		log.Fatalf("查询用户信息错误!", getUserInfoErr)
		response.GetUserInfoFailed(c, "查询用户信息错误 /user/")
		return
	}

	response.GetUserInfoSuccess(c, userId, query_user_id, user)
}

func Register(c *gin.Context) {

}
