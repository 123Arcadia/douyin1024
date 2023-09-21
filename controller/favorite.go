package controller

import (
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/initConfig"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/models"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/response"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

var FavoriteService service.FavoriteService

// FavoriteAction 用户对视频的点赞和取消点赞操
// Params: token/video_id/action_type, 没有user_id
// 1-点赞，2-取消点赞
func FavoriteAction(c *gin.Context) {
	//video_id=1&action_type=1
	//token := c.Query("token")
	videoId, _ := strconv.ParseUint(c.Query("video_id"), 10, 64)
	actionType := c.Query("action_type")

	userId := c.Value("user_id").(uint)
	//fmt.Println("c.Keys =", c.Keys)
	//c.Keys = map[user_id:8]
	switch actionType {
	case "1":
		err := FavoriteService.AddLike(userId, uint(videoId))
		if err != nil {
			// 点赞失败
			response.LikeActionResponse(c, 1, "点赞失败")
			return
		}
		// 点赞成功
		response.LikeActionResponse(c, 0, "点赞成功")
		return
	case "2":
		err := FavoriteService.CancelLike(userId, uint(videoId))
		if err != nil {
			// 点赞失败
			response.LikeActionResponse(c, 1, "点赞失败")
			return
		}
		// 点赞成功
		response.LikeActionResponse(c, 0, "点赞成功")
		return
	default:
		// 异常
		response.FavoriteOperateResponse(c, "，操作失败")
	}
}

// FavoriteList 用户的所有点赞视频
func FavoriteList(c *gin.Context) {
	userId := c.GetUint64("user_id")

	videoIds, err := FavoriteService.GetFavoriteVideoIdList(uint(userId))

	if err != nil {
		response.FavoriteOperateResponse(c, "，操作失败")
	}
	// 获取是有的video
	videoModelsList, err := VideoService.GetVideoByVideoIds(videoIds)
	if err != nil {
		response.FavoriteOperateResponse(c, "查询视频组信息失败")
	}
	videoListResponse := make(response.VideoList, 0, len(videoModelsList))
	for _, video := range videoModelsList {
		videoListResponse = append(videoListResponse, response.Video{
			Id: int64(video.ID),
			Author: response.User{
				Id:             int64(video.User.ID),
				Name:           video.User.UserName,
				FollowCount:    int64(video.User.FollowCount),
				FollowerCount:  int64(video.User.FollowerCount),
				IsFollow:       models.IsFollow(uint(userId), video.User.ID),
				Avatar:         video.User.Avatar,
				Background:     video.User.BackgroundImage,
				Signature:      video.User.Signature,
				TotalFavorited: video.User.TotalFavorited,
				WorkCount:      video.User.WorkCount,
				FavoriteCount:  video.User.FavoriteCount,
			},
			PlayUrl:       initConfig.ScourceStaticPublicPath + video.PlayUrl,
			CoverUrl:      initConfig.ScourceStaticPublicPath + video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    models.IsFavorite(uint(userId), video.ID),
			Title:         video.Description,
		})
	}
	response.FavoriteListResponse(c, videoListResponse)
}
