package models

import (
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/conf"
	"gorm.io/gorm"
	"time"
)

// 评论
type Comment struct {
	ID        uint      `gorm:"primarykey; comment:评论id"`
	UserId    uint      `gorm:"not null;   comment:发布评论的用户id;  type:INT"`
	VideoId   uint      `gorm:"not null;   comment:评论所属视频id;    type:INT"`
	Content   string    `gorm:"not null;   comment:评论内容;          type:VARCHAR(255)"`
	CreatedAt time.Time `gorm:"not null;   comment:评论发布日期;      type:DATETIME"`
	// 定义外键关系
	User  User  `gorm:"foreignKey:UserId; references:ID; comment:评论所属用户"`
	Video Video `gorm:"foreignKey:VideoId; references:ID; comment:评论所属视频"`
}

// IncrementCommentCount 增加该视频的评论总数
func IncrementCommentCount(videoId uint) error {
	res := conf.DB.Model(&Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count + ?", 1))
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// DecreaseCommentCount 删除该视频的评论总数
func DecreaseCommentCount(videoId uint) error {
	res := conf.DB.Model(&Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count - ?", 1))
	if res.Error != nil {
		return res.Error
	}
	return nil
}
