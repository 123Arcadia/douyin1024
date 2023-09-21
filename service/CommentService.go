package service

import (
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/conf"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/initConfig"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/models"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/response"
	"log"
	"time"
)

type CommentService struct {
}

// DeleteComment 删除评论
func (s *CommentService) DeleteComment(userId uint, videoId uint, commentId uint) error {
	err := conf.DB.Where("user_id=? and video_id=? and id=?", userId, videoId, commentId).Delete(&models.Comment{}).Error
	if err != nil {
		return err
	}
	// 删除对应视频的评论总数
	if err = models.DecreaseCommentCount(uint(videoId)); err != nil {
		log.Fatalf("删除该视频的评论总数失败 %+v\n", err)
		return err
	}
	return nil
}

// CreateComment 发布新评论
func (s *CommentService) CreateComment(videoId int64, userId uint, commentText string) (response.Comment, error) {
	// 新建评论
	comment := models.Comment{
		UserId:    userId,
		VideoId:   uint(videoId),
		Content:   commentText,
		CreatedAt: time.Now(), // 差User、Video, 是外键
	}
	err := conf.DB.Create(&comment).Error
	if err != nil {
		log.Fatalf("创建新评论失败 %+v\n", err)
		return response.Comment{}, err
	}
	// 添加该视频的评论总数
	if err = models.IncrementCommentCount(uint(videoId)); err != nil {
		log.Fatalf("添加该视频的评论总数失败 %+v\n", err)
		return response.Comment{}, err
	}

	// 根据user_id获取用户
	user, _ := models.GetUserInfoByUserId(int64(userId))
	// 根据videoId获取对应videoUser
	videoUser, _ := models.GetVideoUserByVideoId(uint(videoId))
	// 评论成功返回评论内容，不需要重新拉取整个列表
	commentResp := response.Comment{
		Id: int64(comment.ID),
		User: response.User{
			Id:             int64(user.ID),
			Name:           user.UserName,
			FollowCount:    int64(user.FollowCount),
			FollowerCount:  int64(user.FollowerCount),
			IsFollow:       models.IsFollow(userId, videoUser.UserId),
			Avatar:         user.Avatar,
			Background:     user.BackgroundImage,
			Signature:      user.Signature,
			FavoriteCount:  user.FavoriteCount,
			TotalFavorited: user.TotalFavorited,
			WorkCount:      user.WorkCount,
		},
		Content:    commentText,
		CreateDate: time.Now().Format(initConfig.TimeFormat),
	}
	return commentResp, nil
}

// GetAllCommentByVideoId 得到视频所有的评论
func (s *CommentService) GetAllCommentByVideoId(videoId uint) ([]models.Comment, error) {
	var commentList []models.Comment
	if err := conf.DB.Where("video_id = ?", videoId).Find(&commentList).Error; err != nil {
		return commentList, err
	}
	return commentList, nil
}
