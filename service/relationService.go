package service

import (
	"errors"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/conf"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/models"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/utils"
	"gorm.io/gorm"
)

type RelationService struct {
}

// GetFriendUser 获取好友列表
func (s *RelationService) GetFriendUser(userId uint) ([]models.User, error) {
	// 通过relation表查询用户关注的好友
	var fromUserIds []uint // 关注该用户的
	var toUserIds []uint   // 该用户关注的
	var allUserIds []uint
	var allUsers []models.User
	if err := conf.DB.Model(&models.Relation{}).Select("to_user_id").Where("from_user_id = ? AND cancel = ?", userId, 0).Find(&fromUserIds).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return allUsers, err
		}
	}
	if err := conf.DB.Model(&models.Relation{}).Select("from_user_id").Where("to_user_id = ? AND cancel = ?", userId, 0).Find(&toUserIds).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return allUsers, err
		}
	}
	// 把fromUserIds、toUserIds的元素结合到allUserIds,并去重
	allUserIds = fromUserIds
	allUserIds = append(allUserIds, toUserIds...)
	allUserIds = utils.RemoveRepeatArray(allUserIds)
	// 获取user信息
	for _, id := range allUserIds {
		user, err := models.GetUserInfoByUserId(int64(id))
		if err != nil {
			continue
		}
		allUsers = append(allUsers, user)
	}
	return allUsers, nil
}

// GetFollowerLists 粉丝列表
func (s *RelationService) GetFollowerLists(userId uint) ([]models.User, error) {
	// 查询该userId的所有关注的userIds
	var relaUserIds []uint
	// 拉去粉丝列表(别人关注该用户, 即to_user_id = userId, cancel = 0)
	err := conf.DB.Model(&models.Relation{}).Select("from_user_id").Where("to_user_id = ? AND cancel = ?", userId, 0).Find(&relaUserIds).Error
	if err != nil {
		return nil, err
	}
	// 根据ids查询Users
	var users []models.User
	for _, id := range relaUserIds {
		curUser, err := models.GetUserInfoByUserId(int64(id))
		if err != nil {
			continue
		}
		users = append(users, curUser)
	}
	return users, nil
}

// GetFollowLists 关注了列表
func (s *RelationService) GetFollowLists(userId uint) ([]models.User, error) {
	var relaUsers []models.User
	var toUserIds []uint
	if err := conf.DB.Model(&models.Relation{}).Select("to_user_id").Where("from_user_id = ?", userId).Find(&toUserIds).Error; err != nil {
		return relaUsers, err
	}

	// 有toUserIds获取User列表
	for _, id := range toUserIds {
		user, err := models.GetUserInfoByUserId(int64(id))
		if err != nil {
			continue
		}
		relaUsers = append(relaUsers, user)
	}
	return relaUsers, nil
}

// CancelFollowUser 取关
func (s *RelationService) CancelFollowUser(formUserId uint, toUserId uint) error {
	// 先检查用户手否存在
	_, err := models.GetUserInfoByUserId(int64(formUserId))
	if err != nil {
		return errors.New("该当前用户不存在")
	}
	_, err = models.GetUserInfoByUserId(int64(toUserId))
	if err != nil {
		return errors.New("该所关注用户不存在")
	}
	// 之前是否已有关注过
	var existedRelation models.Relation
	err = conf.DB.Where("from_user_id = ? AND to_user_id = ?", formUserId, toUserId).First(&existedRelation).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("没有关注关系")
		}
		// 不是没有记录，而是其他错误
		return err
	}
	// cancel: 关注为0，取消关注为1
	// 有记录,但不确定cancel
	if existedRelation.Cancel == 1 {
		return errors.New("已经取关过")
	}

	// 更新取关信息
	if err = conf.DB.Model(&existedRelation).Update("cancel", 1).Error; err != nil {
		return err
	}
	// 更新该当前用户的FollowCount
	err = conf.DB.Model(&models.User{}).Where("id = ?", formUserId).UpdateColumn("follow_count", gorm.Expr("follow_count - ?", 1)).Error
	if err != nil {
		return err
	}
	// 更新被关注用户的 FollowerCount
	err = conf.DB.Model(&models.User{}).Where("id = ?", toUserId).UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

// FollowUser 关注操作
func (s *RelationService) FollowUser(formUserId uint, toUserId uint) error {
	// 先检查用户手否存在
	_, err := models.GetUserInfoByUserId(int64(formUserId))
	if err != nil {
		return errors.New("该当前用户不存在")
	}
	_, err = models.GetUserInfoByUserId(int64(toUserId))
	if err != nil {
		return errors.New("该所关注用户不存在")
	}
	// 之前是否已有关注过
	var existedRelation models.Relation
	err = conf.DB.Where("from_user_id = ? AND to_user_id = ?", formUserId, toUserId).First(&existedRelation).Error
	recordNotFound := false
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			recordNotFound = true
		} else {
			// 不是没有记录，而是其他错误
			return err
		}
	}
	// cancel: 关注为0，取消关注为1
	// 有记录,但不确定cancel
	if !recordNotFound && existedRelation.Cancel == 0 {
		// 已经关注过
		return errors.New("已经关注过")
	}

	if recordNotFound {
		// 需要新建关注关系
		relation := models.Relation{
			FromUserId: formUserId,
			ToUserId:   toUserId,
			Cancel:     0,
		}
		// 保存到DB
		if err = conf.DB.Create(&relation).Error; err != nil {
			return err
		}
		// 更新该当前用户的FollowCount
		err = conf.DB.Model(&models.User{}).Where("id = ?", formUserId).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1)).Error
		if err != nil {
			return err
		}
		// 更新被关注用户的 FollowerCount
		err = conf.DB.Model(&models.User{}).Where("id = ?", toUserId).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error
		if err != nil {
			return err
		}

	}
	return nil
}
