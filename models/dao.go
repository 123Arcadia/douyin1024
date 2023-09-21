package models

import (
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/conf"
	"gorm.io/gorm"
)

// 判断登录用户是否关注了视频作者
func IsFollow(fromUserId, toUserId uint) bool {
	// 查找存在的关注关系
	var existingRelation Relation

	if err := conf.DB.Where("from_user_id = ? AND to_user_id = ? AND cancel = 0", fromUserId, toUserId).First(&existingRelation).Error; err != nil {
		return false
	}
	return true
}

// IsFavorite 判断用户是否点赞当前视频
func IsFavorite(userId, videoId uint) bool {
	// 创建一个 Favorite 结构体实例，用于存储查询结果
	var favorite Favorite
	// 在数据库中查找匹配的点赞记录
	result := conf.DB.Where("user_id = ? AND video_id = ?", userId, videoId).First(&favorite)
	// 检查是否找到匹配的点赞记录
	return result.Error == nil
}

func GetUserInfoByUserId(userId int64) (User, error) {
	var user User
	err := conf.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func GetVideoUserByVideoId(videoId uint) (Video, error) {
	var video Video
	err := conf.DB.Where("id = ?", videoId).First(&video).Error
	if err != nil {
		return Video{}, err
	}
	return video, nil
}
func GetUserVideoCountByUserId(userId uint) (int64, error) {
	var num int64
	err := conf.DB.Model(&Video{}).Where("user_id = ?", userId).Count(&num).Error
	if err != nil {
		return num, err
	}
	return num, nil
}

func IncrementVideoLikeCount(videoId uint) error {
	tx := conf.DB.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1))
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// IncrementUserLikeCount 用户的点赞树增加+1
func IncrementUserLikeCount(userID uint) error {
	return conf.DB.Model(&User{}).Where("id = ?", userID).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
}

// GetAuthorIDForVideo 根据视频ID获取视频的作者ID
func GetAuthorIDForVideo(videoId uint) (uint, error) {
	var userId uint
	err := conf.DB.Model(&Video{}).Select("user_id").Where("id = ?", videoId).Scan(&userId).Error
	if err != nil {
		return 0, err

	}
	return userId, nil
}

// 将视频作者的获赞总数加一
func IncrementAuthorTotalFavorited(author_id uint) error {
	return conf.DB.Model(&User{}).Where("id = ?", author_id).Update("total_favorited", gorm.Expr("total_favorited + ?", 1)).Error
}

// 视频获赞数减一
func DecrementVideoLikeCount(videoId uint) error {
	tx := conf.DB.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1))
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// 取消用户的点赞总数
func DecrementUserLikeCount(userId uint) error {
	err := conf.DB.Model(&User{}).Where("id = ?", userId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
	if err != nil {
		return err

	}
	return nil
}

// 取消的视频作者获赞数减一
func DecrementAuthorTotalFavorited(userId uint) error {
	return conf.DB.Model(&User{}).Where("id = ?", userId).Update("total_favorited", gorm.Expr("total_favorited - ?", 1)).Error
}
