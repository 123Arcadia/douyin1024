package service

import (
	"fmt"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/conf"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/models"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/utils"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path/filepath"
)

type VideoService struct {
}

// 有videoids获取一组视频
func (videoService VideoService) GetVideoByVideoIds(videoIds []uint) (videoList []models.Video, err error) {
	err = conf.DB.Where("id IN ?", videoIds).Preload("User").Find(&videoList).Error
	return videoList, err
}

func (videoService VideoService) GetUserPublishListByUserId(userId uint64) ([]*models.Video, error) {
	var videoList []*models.Video
	err := conf.DB.Where("user_id = ?", userId).Preload("User").Find(&videoList).Error
	if err != nil {
		return videoList, err
	}
	return videoList, nil
}

func (videoService VideoService) Feed(time string) *[]models.Video {
	var videoList *[]models.Video
	conf.DB.Where("created_at <= ?", time).Preload("User").
		Order("created_at DESC").Limit(30).Find(&videoList)
	return videoList
}

// SaveVideoFile 保存文件
func (videoService VideoService) SaveVideoFile(c *gin.Context, file *multipart.FileHeader, title string, userId uint) error {
	// 得到文件存放路径 ./public/videos/
	// 得到该文件在作者发布的第几个顺序
	userVideoCount, err := models.GetUserVideoCountByUserId(userId)
	if err != nil {
		return err
	}
	fileDst := utils.GetVideoNewName(userId, userVideoCount+1, file.Filename, title)
	dst := filepath.Join("./public/videos/", fileDst)
	fmt.Println("视频存储文件路径:", dst)
	if err = c.SaveUploadedFile(file, dst); err != nil {

		return err

	}
	return nil
}
