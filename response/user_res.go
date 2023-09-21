package response

import (
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Id             int64  `json:"id,omitempty"`               // 用户id
	Name           string `json:"name,omitempty"`             // 用户名称
	FollowCount    int64  `json:"follow_count,omitempty"`     // 关注总数
	FollowerCount  int64  `json:"follower_count,omitempty"`   // 粉丝总数
	IsFollow       bool   `json:"is_follow,omitempty"`        // true-已关注，false-未关注
	Avatar         string `json:"avatar,omitempty"`           // 用户头像
	Background     string `json:"background_image,omitempty"` // 用户个人页顶部大图
	Signature      string `json:"signature,omitempty"`        // 个人简介
	FavoriteCount  int64  `json:"favorite_count"`             // 喜欢数
	TotalFavorited string `json:"total_favorited"`            // 获赞数量
	WorkCount      int64  `json:"work_count"`                 // 作品数
}

// 用户注册响应
type UserRegisterResponse struct {
	StatusCode int32  `json:"status_code"`       // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`        // 返回状态描述
	UserId     int32  `json:"user_id,omitempty"` // 用户鉴权token
	Token      string `json:"token,omitempty"`   // 用户id
}

// 用户登录响应
type UserLoginResponse struct {
	StatusCode int32  `json:"status_code"`       // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`        // 返回状态描述
	UserId     int32  `json:"user_id,omitempty"` // 用户鉴权token
	Token      string `json:"token,omitempty"`   // 用户id
}

// 用户信息响应
type UserInfoResponse struct {
	StatusCode int32  `json:"status_code"`    // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`     // 返回状态描述
	User       User   `json:"user,omitempty"` // 用户信息
}

// GetUserInfoFailed 得到用户信息
func GetUserInfoFailed(c *gin.Context, msg string) {
	c.JSON(500, UserInfoResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}

func GetUserInfoSuccess(c *gin.Context, userId uint, query_user_id uint64, user models.User) {

	c.JSON(http.StatusOK, UserInfoResponse{
		StatusCode: 0,
		StatusMsg:  "用户信息获取成功",
		User: User{
			Id:             int64(user.ID),
			Name:           user.UserName,
			FollowCount:    int64(user.FollowCount),
			FollowerCount:  int64(user.FollowerCount),
			IsFollow:       models.IsFollow(userId, uint(query_user_id)),
			Avatar:         user.Avatar,
			Background:     user.BackgroundImage,
			Signature:      user.Signature,
			TotalFavorited: user.TotalFavorited,
			WorkCount:      user.WorkCount,
			FavoriteCount:  user.FavoriteCount,
		},
	})
}

func RegisterUserFailed(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, UserRegisterResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}

func TokenGenernateFailed(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, UserRegisterResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}

func RegisterUserSuccess(c *gin.Context, userId int32, token string) {
	c.JSON(http.StatusOK, UserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "用户注册成功",
		UserId:     int32(userId),
		Token:      token,
	})
}

func UserLoginFailed(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, UserLoginResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}

func UserLoginSuccess(c *gin.Context, userId int32, token string) {
	c.JSON(http.StatusOK, UserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "用户登录成功",
		UserId:     userId,
		Token:      token,
	})
}
