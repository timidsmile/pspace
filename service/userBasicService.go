package service

import (
	"errors"
	"sync"

	"github.com/timidsmile/pspace/components"
	"github.com/timidsmile/pspace/model"
)

type UserBasicService struct {
	mutex *sync.Mutex
}

func (s *UserBasicService) Register(basic model.UserBasic) error {
	// check 该用户是否已经注册过 & 用户ID是否重复
	if hasUserName, hasUserID := s.GetByUserName(basic.UserName), s.GetByUserID(basic.UserID); hasUserName != nil || hasUserID != nil {
		return errors.New("该用户已注册")
	}

	s.CreateUser(&basic)

	return nil
}

func (s *UserBasicService) GetByUserName(userName string) *model.UserBasic {
	userBasic := &model.UserBasic{}
	if err := components.Db.Where("`user_name` = ?", userName).First(userBasic).Error; nil != err {
		return nil
	}

	return userBasic
}

func (u *UserBasicService) CreateUser(user *model.UserBasic) error {
	components.Db.Create(&user)
	return nil
}

// 根据userID查询userBasic信息
func (srv *UserBasicService) GetByUserID(userID int64) *model.UserBasic {
	userBasic := &model.UserBasic{}
	if err := components.Db.Where("`user_id` = ?", userID).First(userBasic).Error; nil != err {

		return nil
	}

	return userBasic
}

// 批量查询方式
func (srv *UserBasicService) GetByUserIds(userIds []string) *model.UserBasic {
	userBasic := &model.UserBasic{}
	if err := components.Db.Where("`user_id` in (?)", userIds).Find(userBasic).Error; nil != err {

		return nil
	}

	return userBasic
}

func (srv *UserBasicService) GetByEmail(email string) *model.UserBasic {
	userBasic := &model.UserBasic{}
	if err := components.Db.Where("`email` = ?", email).First(userBasic).Error; nil != err {
		return nil
	}

	return userBasic
}

func (srv *UserBasicService) GetByMobile(mobile string) *model.UserBasic {
	userBasic := &model.UserBasic{}
	if err := components.Db.Where("`mobile` = ?", mobile).First(userBasic).Error; nil != err {
		return nil
	}

	return userBasic
}

// 根据userName前缀匹配查询符合条件的所有记录
func (srv *UserBasicService) GetByUserNamePrefix(userNamePreFix string) *model.UserBasic {
	userBasic := &model.UserBasic{}
	if err := components.Db.Where("`user_name` ilike %?%", userNamePreFix).Find(userBasic).Error; nil != err {

		return nil
	}

	return userBasic
}

// 结构体查询方式（多个匹配条件时）
func (srv *UserBasicService) GetByCondition(user model.UserBasic) *model.UserBasic {
	userBasic := &model.UserBasic{}
	if err := components.Db.Where(&user).Find(userBasic).Error; nil != err {

		return nil
	}

	return userBasic
}

// or 查询方式
// TODO：or查询方式
func (srv *UserBasicService) GetByID(userName string, email string, mobile string) *model.UserBasic {
	userBasic := &model.UserBasic{}
	if err := components.Db.Where("user_name = ?", userName).Or("email = ?", email).Or("mobile = ?", mobile).Find(userBasic).Error; nil != err {

		return nil
	}

	return userBasic
}

func (s *UserBasicService) UserSetting(updateBasic model.UserBasic) error {
	// check 该用户是否已经注册过

	userID := updateBasic.UserID
	users := s.GetByUserID(userID)

	if users == nil {
		return errors.New("用户未注册！")
	}

	return components.Db.Table(updateBasic.TableName()).Where("user_id = ?", userID).Updates(updateBasic).Error
}
