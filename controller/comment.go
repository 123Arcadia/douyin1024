package controller

import (
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/initConfig"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/models"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/response"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

var CommentService service.CommentService

// CommentAction 登录用户对视频进行评论
// 参数：token、video_id、action_type(1-发布评论，2-删除评论)、
// comment_text(在action_type=1的时候使用)、comment_id(在action_type=2的时候使用)
func CommentAction(c *gin.Context) {
	// 从上下文中获取user_id
	userId := c.GetUint("user_id")
	video_id, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		// video_id类型转换失败
		response.VideoIdConversionErr(c)
	}
	action_type := c.Query("action_type")
	switch action_type {
	case "1":
		// 使用comment_text
		comment_text := c.Query("comment_text")
		// 发布评论
		commentResp, err := CommentService.CreateComment(video_id, userId, comment_text)
		if err != nil {
			response.CommentActionResponseHandler(c, 1, "评论保存失败", commentResp)
		}
		// 评论添加成功
		response.CommentActionResponseHandler(c, 0, "评论添加成功", commentResp)
		log.Println("[CommentAction]", c.Params, "-", c.Keys, "-", action_type, "-")

	case "2":
		// 使用comment_id， 删除评论
		comment_id, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		if err != nil {
			// comment_id类型转换失败
			response.CommentIdConversionErr(c)
		}
		// 删除评论
		err = CommentService.DeleteComment(userId, uint(video_id), uint(comment_id))
		if err != nil {
			response.CommentActionResponseDelHandler(c, 1, "评论删除失败")
		}
		// 评论添加成功
		response.CommentActionResponseDelHandler(c, 0, "评论删除成功")
		log.Println("[CommentAction]", c.Params, "-", c.Keys, "-", action_type, "-")
		//2023/09/21 01:27:36 [CommentAction] [] - map[user_id:8] - 1 -
	default:
		// 操作异常
		response.CommentOperateResponse(c)
	}
}

// CommentList 查看视频的所有评论，按发布时间倒序
func CommentList(c *gin.Context) {
	videoId, err := strconv.ParseUint(c.Query("video_id"), 10, 64)
	if err != nil {
		// video_id类型转换失败
		response.VideoIdConversionErr(c)
	}
	// 查询该视频所有评论
	commenLists, err := CommentService.GetAllCommentByVideoId(uint(videoId))
	if err != nil {
		// 评论列表获取失败
		response.GetCommentListResponseHandler(c, 1, "评论列表获取失败", nil)
		return
	}
	// 由上下文获取user_id
	userId := c.GetUint("user_id")

	// 以comment_resp返回
	commentListResp := make([]response.Comment, 0, len(commenLists))
	for _, comment := range commenLists {
		// 由comment.UserId获取User
		user, GetUserInfoByUserIdErr := models.GetUserInfoByUserId(int64(comment.UserId))
		if GetUserInfoByUserIdErr != nil {
			continue
		}
		commentListResp = append(commentListResp, response.Comment{
			Id: int64(comment.ID),
			User: response.User{
				Id:             int64(user.ID),
				Name:           user.UserName,
				FollowCount:    int64(user.FollowCount),
				FollowerCount:  int64(user.FollowerCount),
				IsFollow:       models.IsFollow(userId, user.ID),
				Avatar:         user.Avatar,
				Background:     user.BackgroundImage,
				Signature:      user.Signature,
				TotalFavorited: user.TotalFavorited,
				WorkCount:      user.WorkCount,
				FavoriteCount:  user.FavoriteCount,
			},
			Content:    comment.Content,
			CreateDate: time.Now().Format(initConfig.TimeFormat),
		})
	}

	// 评论列表获取成功
	log.Println("评论列表获取成功", commentListResp)
	response.GetCommentListResponseHandler(c, 0, "评论列表获取成功", commentListResp)
}
