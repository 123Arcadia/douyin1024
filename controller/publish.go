package controller

import (
	"fmt"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/initConfig"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/models"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Publish data/token/title
func Publish(c *gin.Context) {
	// 视频信息
	// TODO: 2023/9/18 改成支持多文件上传
	file, err := c.FormFile("data")
	// NOTE: 文档中没有写传user_id
	userId, _ := strconv.ParseUint(c.Query("user_id"), 10, 64)
	// 2022-10-09 17-14-26.mp4  size: 24100747
	if err != nil {
		response.VideoFileAccessErr(c, "视频文件获取失败") // 视屏获取失败
		return
	}
	title := c.PostForm("title")
	//-----------------------
	fmt.Println("key:", c.Keys)
	//key: map[user_id:7]
	fmt.Println("params:", c.Params)
	//params: []
	fmt.Println(title, file.Filename, file.Size, "userId =", userId)
	//test3 2022-11-07 18-51-10海尔 (1) (1).mp4 4455537
	//-----------------------
	err = VideoService.SaveVideoFile(c, file, title, uint(userId))
	if err != nil {
		response.SaveVideoFileResponse(c, 1, "视频文件保存失败")
	}
	response.SaveVideoFileResponse(c, 0, "视频文件保存成功")
}

// PublishList 用户发布视频列表
func PublishList(c *gin.Context) {
	fmt.Println("[/douyin/publish/list/]key:", c.Keys, c.Params, c.Request.Body)
	//[/douyin/publish/list/]key: map[user_id:8] [] {}

	//userId := c.GetUint64("user_id")
	//user_id =  1
	query_user_id, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	fmt.Println("query_user_id = ", query_user_id)
	//query_user_id =  1

	videoList, err := VideoService.GetUserPublishListByUserId(query_user_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.VideoListResponse{
			StatusCode: 1,
			StatusMsg:  "视频列表获取失败",
		})
	}
	// 创建videoList的Response
	videoListResponse := make(response.VideoList, 0, len(videoList))
	for _, video := range videoList {
		videoListResponse = append(videoListResponse, response.Video{
			Id: int64(video.ID),
			Author: response.User{
				Id:             int64(video.UserId),
				Name:           video.User.UserName,
				FollowCount:    int64(video.User.FollowCount),
				FollowerCount:  int64(video.User.FollowerCount),
				IsFollow:       models.IsFollow(uint(query_user_id), video.UserId),
				Avatar:         video.User.Avatar,
				Background:     video.User.BackgroundImage,
				Signature:      video.User.Signature,
				TotalFavorited: video.User.TotalFavorited,
				WorkCount:      video.User.WorkCount,
				FavoriteCount:  video.User.FavoriteCount,
			},
			CommentCount:  video.CommentCount,
			CoverUrl:      initConfig.ScourceStaticPublicPath + video.CoverUrl,
			PlayUrl:       initConfig.ScourceStaticPublicPath + video.PlayUrl,
			FavoriteCount: video.FavoriteCount,
			IsFavorite:    models.IsFavorite(uint(query_user_id), video.UserId), // 这里是userId是否对他曾静发过的该video点赞过
			Title:         video.Description,
		})
	}
	c.JSON(http.StatusOK, response.VideoListResponse{
		StatusCode: 0,
		StatusMsg:  "获取视频发布列表成功",
		VideoList:  videoListResponse,
	})
}
