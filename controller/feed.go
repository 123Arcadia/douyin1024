package controller

import (
	"fmt"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/initConfig"
	middlewares "github.com/123Arcadia/douyin1024CodeSpaceDemo.git/middlewares"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/models"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/response"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/service"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

var VideoService service.VideoService

// 不限制登录状态，返回按 投稿时间 倒序 的视频列表，视频数由服务端控制，单次最多30个
// Feed same demo video list for every request
func Feed(c *gin.Context) {
	tokenString := c.Query("token")
	if len(tokenString) == 0 {
		// 未登录
		NoLoginAccess(c)
		return
	}
	// 鉴权
	token, err := jwt.ParseWithClaims(tokenString, &middlewares.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return initConfig.AUTH_KEY, nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, middlewares.AuthFailResponse{
			StatusCode: 1,
			StatusMsg:  middlewares.ParseTokenFailed,
		})
		c.Abort()
		return
	}
	// 检验
	claims, ok := token.Claims.(*middlewares.MyClaims)
	if ok && token.Valid {
		fmt.Printf("鉴权完成，已登录 %v %v", claims.UserID, claims.Username)
		LoginAccess(c, claims.UserID)
		return
	} else {
		fmt.Println(err)
	}
	c.JSON(http.StatusInternalServerError, middlewares.AuthFailResponse{
		StatusCode: 1,
		StatusMsg:  middlewares.CheckFailed,
	})
}

func LoginAccess(c *gin.Context, userId uint) {
	//  需要返回的：最新投稿时间戳，不填表示当前时间
	startTime := utils.GetFormatTime(c.Query("last_time"))
	feedVideoList := *VideoService.Feed(startTime)
	lenFeedVideoList := len(feedVideoList)
	if lenFeedVideoList <= 0 {
		log.Fatalf("视频流为空 error_2")
		// 意味着startTime参数没有填写，按默认time.now()来查找30个
		response.GetFeedSuccess(c, time.Now().Unix(), []response.Video{})
		return
	}
	//返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个
	lastTime := feedVideoList[lenFeedVideoList-1].CreatedAt.Unix()
	videoList_res := make([]response.Video, 0, lenFeedVideoList)
	for _, video := range feedVideoList {
		videoList_res = append(videoList_res, response.Video{
			Id: int64(video.ID),
			Author: response.User{
				Id:             int64(video.User.ID),
				Name:           video.User.UserName,
				FollowCount:    int64(video.User.FollowCount),
				FollowerCount:  int64(video.User.FollowerCount),
				IsFollow:       models.IsFollow(userId, video.User.ID), // 登录状态显示:检查该用户对该video作者是否关注
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
			IsFavorite:    models.IsFavorite(userId, video.ID), // 检查该用户对该video是否点赞
			Title:         video.Description,
			// 没有写isFavority
		})
	}
	response.GetFeedSuccess(c, lastTime, videoList_res)
}

// 未登录自动推送视频
func NoLoginAccess(c *gin.Context) {
	//  需要返回的：最新投稿时间戳，不填表示当前时间
	startTime := utils.GetFormatTime(c.Query("last_time"))
	feedVideoList := *VideoService.Feed(startTime)
	lenFeedVideoList := len(feedVideoList)
	if lenFeedVideoList <= 0 {
		log.Fatalf("获取视频流错误! error_1")
		// 意味着startTime参数没有填写，按默认time.now()来查找30个
		response.GetFeedSuccess(c, time.Now().Unix(), []response.Video{})
		return
	}
	//返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个
	lastTime := feedVideoList[lenFeedVideoList-1].CreatedAt.Unix()
	videoList_res := make([]response.Video, 0, lenFeedVideoList)
	for _, video := range feedVideoList {
		videoList_res = append(videoList_res, response.Video{
			Id: int64(video.ID),
			Author: response.User{
				Id:             int64(video.User.ID),
				Name:           video.User.UserName,
				FollowCount:    int64(video.User.FollowCount),
				FollowerCount:  int64(video.User.FollowerCount),
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
			Title:         video.Description,
			// 没有写isFavority: // 登录状态显示
		})
	}
	response.GetFeedSuccess(c, lastTime, videoList_res)
}
