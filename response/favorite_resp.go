package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type FavoriteResponse struct {
	StatusCode int64  `json:"Status_code"`
	StatusMsg  string `json:"status_msg"` // 返回状态描述

}

func LikeActionResponse(c *gin.Context, code int64, msg string) {
	var httpStat int
	if code == 1 {
		httpStat = 500
	} else {
		httpStat = 200
	}
	c.JSON(httpStat, FavoriteResponse{
		StatusCode: code,
		StatusMsg:  msg,
	})
}

func FavoriteOperateResponse(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, FavoriteResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}

// FavoriteList 喜欢列表响应
type FavoriteList struct {
	StatusCode string  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`           // 返回状态描述
	VideoList  []Video `json:"video_list,omitempty"` // 用户点赞视频列表
}

func FavoriteListResponse(c *gin.Context, videoListResponse []Video) {
	c.JSON(http.StatusOK, FavoriteList{
		StatusCode: "0",
		StatusMsg:  "视频列表获取成功",
		VideoList:  videoListResponse,
	})
}
