package service

import (
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/conf"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/models"
	"time"
)

type MessageService struct {
}

// GetMsgListWithTime 获取信息列表
func (s *MessageService) GetMsgListWithTime(userId uint, toUserId uint, msgTime time.Time) ([]models.Message, error) {
	var msgs []models.Message
	err := conf.DB.Where("((from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)) AND create_time > ?", userId, toUserId, toUserId, userId, msgTime).
		Order("create_time").
		Find(&msgs).
		Error
	if err != nil {
		return []models.Message{}, err
	}
	return msgs, nil

}

// AddMsg 发送消息
func (s *MessageService) AddMsg(userId uint, toUserId uint, ctx string) error {
	message := models.Message{
		FromUserID: userId,
		ToUserID:   toUserId,
		Content:    ctx,
		CreateTime: time.Now(),
	}
	// 插入到DB
	err := conf.DB.Create(&message).Error
	if err != nil {
		return err
	}
	return nil
}
