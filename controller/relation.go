package controller

import (
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/models"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/response"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

//type UserListResponse struct {
//	Response
//	UserList []User `json:"user_list"`
//}

var RelationService service.RelationService

// RelationAction 关注操作
// 参数:token/to_user_id/action_type(1-关注，2-取消关注)
func RelationAction(c *gin.Context) {
	// 登录的用户id
	formUserId := c.GetUint("user_id")
	toUerId, err := strconv.ParseUint(c.Query("to_user_id"), 10, 64)
	if err != nil {
		response.ToUserIdConversionErr(c, "对方用户id参数类型转换失败")
		return
	}
	actionType := c.Query("action_type")
	if formUserId == uint(toUerId) {
		response.RelationOperateResponseHandler(c, 1, "操作失败(自己不用关注或取关自己)")
		return
	}
	switch actionType {
	case "1":
		// 关注
		err = RelationService.FollowUser(formUserId, uint(toUerId))
		if err != nil {
			response.RelationOperateResponseHandler(c, 1, "关注失败")
			return
		}
		response.RelationOperateResponseHandler(c, 0, "关注成功")
		return
	case "2":
		err = RelationService.CancelFollowUser(formUserId, uint(toUerId))
		if err != nil {
			response.RelationOperateResponseHandler(c, 1, "取关失败")
			return
		}
		response.RelationOperateResponseHandler(c, 0, "取关成功")
		// 取关
		return
	default:
		response.RelationOperateResponseHandler(c, 1, "关注操作异常")
		return
	}

}

// FollowList 关注列表
func FollowList(c *gin.Context) {
	// 用户id
	userId := c.GetUint("user_id")
	if userId < 0 {
		response.ToUserIdConversionErr(c, "参用户id错误")
	}
	relaUsers, err := RelationService.GetFollowLists(userId)
	if err != nil {
		response.FollowListsResponseHandler(c, "1", "获取关注列表异常", nil)
		return
	}
	//var relaUserResponse []response.User
	relaUserResponse := make([]response.User, 0, len(relaUsers))
	for _, user := range relaUsers {
		relaUserResponse = append(relaUserResponse, response.User{
			Id:             int64(user.ID),
			Name:           user.UserName,
			Avatar:         user.Avatar,
			Signature:      user.Signature,
			FollowCount:    int64(user.FollowCount),
			FollowerCount:  int64(user.FollowCount),
			IsFollow:       models.IsFollow(userId, user.ID),
			Background:     user.BackgroundImage,
			TotalFavorited: user.TotalFavorited,
			WorkCount:      user.WorkCount,
			FavoriteCount:  user.FavoriteCount})
	}
	// 关注成功
	response.FollowListsResponseHandler(c, "0", "获取关注列表成功", relaUserResponse)

}

// FollowerList 粉丝列表, parmas: user_id、token
func FollowerList(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		response.ToUserIdConversionErr(c, "参数类型转换失败")
	}
	// 获取粉丝列表
	follerLists, err := RelationService.GetFollowerLists(uint(userId))
	if err != nil {
		response.FollowerListsResponseHandler(c, "1", "粉丝列表获取失败", nil)
		return
	}
	// 转换为user_resp
	var userResp []response.User
	for _, user := range follerLists {
		userResp = append(userResp, response.User{
			Id:             int64(user.ID),
			Name:           user.UserName,
			Avatar:         user.Avatar,
			Signature:      user.Signature,
			FollowCount:    int64(user.FollowCount),
			FollowerCount:  int64(user.FollowCount),
			IsFollow:       models.IsFollow(uint(userId), user.ID),
			Background:     user.BackgroundImage,
			TotalFavorited: user.TotalFavorited,
			WorkCount:      user.WorkCount,
			FavoriteCount:  user.FavoriteCount,
		})
	}
	response.FollowerListsResponseHandler(c, "0", "粉丝列表获取成功", userResp)

	// FriendList all users have same friend list
	//func FriendList(c *gin.Context) {
	//	c.JSON(http.StatusOK, UserListResponse{
	//		Response: Response{
	//			StatusCode: 0,
	//		},
	//		UserList: []User{DemoUser},
	//	})
}

// FriendList 获取好友列表
func FriendList(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		response.ToUserIdConversionErr(c, "参数类型转换失败")
	}
	// 好友列表
	friendUsers, err := RelationService.GetFriendUser(uint(userId))
	if err != nil {
		response.RelationGetFriendsResponseHandler(c, "1", "获取好友列表失败", nil)
	}

	var friendUsersResp []response.User
	for _, user := range friendUsers {
		friendUsersResp = append(friendUsersResp, response.User{
			Id:             int64(user.ID),
			Name:           user.UserName,
			Avatar:         user.Avatar,
			Signature:      user.Signature,
			FollowCount:    int64(user.FollowCount),
			FollowerCount:  int64(user.FollowCount),
			IsFollow:       models.IsFollow(uint(userId), user.ID),
			Background:     user.BackgroundImage,
			TotalFavorited: user.TotalFavorited,
			WorkCount:      user.WorkCount,
			FavoriteCount:  user.FavoriteCount,
		})
	}
	log.Printf("用户id = %+v 的好友列表 friendUsers = %+v\n", userId, friendUsers)
	response.RelationGetFriendsResponseHandler(c, "0", "获取好友列表成功", friendUsersResp)
}
