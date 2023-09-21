package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentResponse struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}
type Comment struct {
	Id         int64  `json:"id"`          // 评论id
	User       User   `json:"user"`        // 评论用户
	Content    string `json:"content"`     // 评论内容
	CreateDate string `json:"create_date"` // 评论发布日期，格式 mm-dd
}

// 评论操作响应
type CommentActionResponse struct {
	StatusCode int32   `json:"status_code"`       // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`        // 返回状态描述
	Comment    Comment `json:"comment,omitempty"` // 评论成功返回评论内容，不需要重新拉取整个列表
}

func CommentOperateResponse(c *gin.Context) {
	c.JSON(http.StatusBadRequest, CommentActionResponse{
		StatusCode: 1,
		StatusMsg:  "，评论操作失败",
	})
}

func CommentActionResponseHandler(c *gin.Context, code int32, msg string, commentResp Comment) {
	var httpStat int
	if code == 1 {
		httpStat = 400
	} else {
		httpStat = 200
	}
	c.JSON(httpStat, CommentActionResponse{
		StatusCode: code,
		StatusMsg:  msg,
		Comment:    commentResp,
	})
}

// CommentIdConversionErr  comment_id 类型转换
func CommentIdConversionErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, CommentResponse{
		StatusCode: 1,
		StatusMsg:  "评论id类型转换失败",
	})
}

// CommentActionResponseDelHandler 删除评论响应
func CommentActionResponseDelHandler(c *gin.Context, code int32, msg string) {
	var httpStat int
	if code == 1 {
		httpStat = 400
	} else {
		httpStat = 200
	}
	c.JSON(httpStat, CommentActionResponse{
		StatusCode: code,
		StatusMsg:  msg,
	})
}

// 评论列表响应
type CommentListResponse struct {
	StatusCode  int       `json:"status_code"`            // 状态码，0-成功，其他值-失败
	StatusMsg   string    `json:"status_msg"`             // 返回状态描述
	CommentList []Comment `json:"comment_list,omitempty"` // 评论列表
}

// GetCommentListResponseHandler 评论列表获取情况
func GetCommentListResponseHandler(c *gin.Context, code int, msg string, commentLists []Comment) {
	var httpStat int
	if code == 1 {
		httpStat = 400
	} else {
		httpStat = 200
	}
	c.JSON(httpStat, CommentListResponse{
		StatusCode:  code,
		StatusMsg:   msg,
		CommentList: commentLists,
	})
}
