package service

import (
	"errors"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/conf"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/initModelsExamples"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/models"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/utils"
	"gorm.io/gorm"
	"log"
)

type UserService struct {
}

// Register 注册
func (s *UserService) Register(username string, password string) (models.User, error) {
	var user models.User
	err := conf.DB.Where("user_name = ?", username).First(&user).Error
	if err == nil {
		log.Printf("用户名已注册username = %+v\n", username)
		return user, errors.New("用户名已注册")
	}
	// 加密password
	encryptPassword, err := utils.EncryptPassword(password)
	if err != nil {
		log.Printf("用户密码不符username = %+v, password = %+v\n", username, password)
		return models.User{}, err
	}

	user = models.User{
		UserName:        username,
		PassWord:        string(encryptPassword),
		FollowCount:     0,
		FollowerCount:   0,
		FavoriteCount:   0,
		TotalFavorited:  "0",
		WorkCount:       0,
		Avatar:          initModelsExamples.UserDefault_avatar_url,
		BackgroundImage: initModelsExamples.UserDefault_background_image_url,
		Signature:       initModelsExamples.UserDefault_profile_description,
	}
	err = conf.DB.Create(&user).Error
	if err != nil {
		log.Fatalf("用户创建失败 err = %v", err)
		return models.User{}, err
	}
	return user, nil
}

// Login 登录
func (s *UserService) Login(username string, password string) (models.User, error) {
	var user models.User
	err := conf.DB.Where("user_name = ?", username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		log.Printf("用户不存在 %v\n", username)
		return models.User{}, errors.New("用户不存在")
	} else if err != nil {
		log.Printf("登录失败 %v\n", username)
		return models.User{}, err
	}
	// 验证密码
	if utils.CheckPasswordValidity(user.PassWord, password) {
		return user, nil
	} else {
		return models.User{}, errors.New("用户密码错误")
	}
}
