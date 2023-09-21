package service

import (
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/conf"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/models"
	"gorm.io/gorm"
)

type FavoriteService struct {
}

func (s FavoriteService) CancelLike(userId, videoId uint) error {

	var existedFavorite models.Favorite
	//查找之前是有这种点赞关系存在
	err := conf.DB.Where("user_id = ? AND video_id = ?", userId, videoId).First(&existedFavorite).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 点赞关系不存在, 则无反应
			return nil
		}
		return err
	}
	err = conf.DB.Delete(&existedFavorite).Error
	if err != nil {
		return err
	}
	// 用户的点赞数减一
	if err := models.DecrementUserLikeCount(userId); err != nil {
		return err
	}

	// 视频的获赞数加一
	if err := models.DecrementVideoLikeCount(uint(videoId)); err != nil {
		return err
	}

	// 获取该取消赞的视频作者id
	author_id, err := models.GetAuthorIDForVideo(uint(videoId))
	if err != nil {
		return err
	}

	// 将视频作者的获赞总数减一
	if err := models.DecrementAuthorTotalFavorited(author_id); err != nil {
		return err
	}

	return nil
}

func (s *FavoriteService) AddLike(userId, videoId uint) error {
	favorite := models.Favorite{
		UserId:  userId,
		VideoId: videoId,
	}
	// 新建保存一天点赞
	err := conf.DB.Create(&favorite).Error
	if err != nil {
		return err
	}

	// 将该视频的获赞总数加一
	if err := models.IncrementVideoLikeCount(uint(videoId)); err != nil {
		return err
	}

	// 将本次点赞用户的点赞数加一
	if err := models.IncrementUserLikeCount(userId); err != nil {
		return err
	}

	// 根据该视频id来获取其作者id
	author_id, err := models.GetAuthorIDForVideo(uint(videoId))
	if err != nil {
		return err
	}
	// 将视频作者的获赞总数加一
	if err := models.IncrementAuthorTotalFavorited(author_id); err != nil {
		return err
	}
	return nil
}

// GetFavoriteVideoIdList 根据用户ID取出该用户点赞的所有视频ID
func (s *FavoriteService) GetFavoriteVideoIdList(userId uint) ([]uint, error) {
	var videoIDs []uint
	var favorites []models.Favorite
	err := conf.DB.Where("user_id = ?", userId).Find(&favorites).Error
	if err != nil {
		return nil, err
	}
	for _, favor := range favorites {
		videoIDs = append(videoIDs, favor.VideoId)
	}
	return videoIDs, nil
}
