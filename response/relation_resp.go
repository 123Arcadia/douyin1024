package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RelationActionResponse struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

func ToUserIdConversionErr(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, RelationActionResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}

func RelationOperateResponseHandler(c *gin.Context, code int32, msg string) {
	var httpStat int32
	if code == 1 {
		httpStat = http.StatusBadRequest
	} else if code == 0 {
		httpStat = http.StatusOK
	}
	c.JSON(int(httpStat), RelationActionResponse{
		StatusCode: code,
		StatusMsg:  msg,
	})

}

// 关注列表响应
type FollowListResponse struct {
	StatusCode string `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	UserList   []User `json:"user_list"`   // 用户信息列表
}

// FollowListsResponseHandler 获取关注列表
func FollowListsResponseHandler(c *gin.Context, code string, msg string, userList []User) {
	var httpStat int
	if code == "0" {
		httpStat = http.StatusOK
	} else {
		httpStat = http.StatusBadRequest

	}
	c.JSON(httpStat, FollowListResponse{
		StatusCode: code,
		StatusMsg:  msg,
		UserList:   userList,
	})
}

// FollowListResponse 粉丝列表获取
type FollowerListResponse struct {
	StatusCode string `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	UserList   []User `json:"user_list"`   // 用户信息列表
}

// FollowerListsResponseHandler 粉丝列表获取
func FollowerListsResponseHandler(c *gin.Context, code string, msg string, userLists []User) {
	var httpStat int
	if code == "0" {
		httpStat = http.StatusOK
	} else {
		httpStat = http.StatusBadRequest

	}
	c.JSON(httpStat, FollowerListResponse{
		StatusCode: code,
		StatusMsg:  msg,
		UserList:   userLists,
	})
}

// 好友列表响应
type FriendListResponse struct {
	StatusCode string `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	UserList   []User `json:"user_list"`   // 用户列表
}

// RelationGetFriendsResponseHandler 获取好友列表
func RelationGetFriendsResponseHandler(c *gin.Context, code string, msg string, userLists []User) {
	var httpStat int
	if code == "0" {
		httpStat = http.StatusOK
	} else {
		httpStat = http.StatusBadRequest
	}
	c.JSON(httpStat, FriendListResponse{
		StatusCode: code,
		StatusMsg:  msg,
		UserList:   userLists,
	})
}
