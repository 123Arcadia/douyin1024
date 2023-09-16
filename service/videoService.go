package service

import (
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/conf"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/models"
)

type VideoService struct {
}

func (s *VideoService) Feed(time string) *[]models.Video {
	var videoList *[]models.Video
	conf.DB.Where("created_at <= ?", time).Preload("User").
		Order("created_at DESC").Limit(30).Find(&videoList)
	return videoList
}
