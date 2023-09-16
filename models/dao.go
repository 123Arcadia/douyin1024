package models

import (
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/conf"
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

// 判断用户是否点赞当前视频
func IsFavorite(userId, videoId uint) bool {
	// 创建一个 Favorite 结构体实例，用于存储查询结果
	var favorite Favorite
	// 在数据库中查找匹配的点赞记录
	result := conf.DB.Where("user_id = ? AND video_id = ?", userId, videoId).First(&favorite)
	// 检查是否找到匹配的点赞记录
	return result.Error == nil
}
