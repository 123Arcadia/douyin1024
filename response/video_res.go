package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Video struct {
	Id            int64  `json:"id,omitempty"`   // 视频唯一标识
	Author        User   `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverUrl      string `json:"cover_url"`      // 视频封面地址
	PlayUrl       string `json:"play_url"`       // 视频播放地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	Title         string `json:"title"`          // 视频标题
}

type VideoList []Video

// 视频流响应
type FeedResponse struct {
	StatusCode int32   `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`           // 返回状态描述
	VideoList  []Video `json:"video_list,omitempty"` // 视频猎豹
	NextTime   int64   `json:"next_time,omitempty"`  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

func GetFeedSuccess(c *gin.Context, lastTime int64, videoFeedList []Video) {
	c.JSON(http.StatusOK, FeedResponse{
		StatusCode: 0,
		StatusMsg:  "获取视频流成功",
		VideoList:  videoFeedList,
		NextTime:   lastTime,
	})
}

type PostVideoResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func VideoFileAccessErr(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, PostVideoResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}

func SaveVideoFileResponse(c *gin.Context, code int32, msg string) {
	var httpStat int
	if code == 1 {
		httpStat = 500
	} else {
		httpStat = 200
	}
	c.JSON(httpStat, PostVideoResponse{
		StatusCode: code,
		StatusMsg:  msg,
	})
}

// 发布列表响应
type VideoListResponse struct {
	StatusCode int32   `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`           // 返回状态描述
	VideoList  []Video `json:"video_list,omitempty"` // 用户发布的视频列表
}

// VideoIdConversionErr 视频id类型转换失败
func VideoIdConversionErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, PostVideoResponse{
		StatusCode: 1,
		StatusMsg:  "视频id类型转换失败",
	})
}