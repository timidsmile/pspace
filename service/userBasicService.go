package service

import (
	"errors"
	"fmt"
	"github.com/timidsmile/pspace/components"
	"github.com/timidsmile/pspace/model"
	"sync"
)

type UserBasicService struct {
	mutex *sync.Mutex
}

func (s *UserBasicService) RegisterByEmail(email string, passwd string, userID int64) error {
	// check 该用户是否已经注册过

	users := s.GetByEmail(email)

	if users != nil {
		return errors.New("该用户已注册")
	}

	user := model.UserBasic{
		UserID:   userID,
		UserName: "",
		Mobile:   "",
		Email:    email,
		Passwd:   passwd,
	}

	fmt.Println("2222")

	s.CreateUser(&user)

	return nil
}

func (s *UserBasicService) RegisterByMobile(mobile string, passwd string, userID int64) error {
	// check 该用户是否已经注册过
	users := s.GetByEmail(mobile)

	if users != nil {
		return errors.New("该用户已注册")
	}

	user := model.UserBasic{
		UserID:   userID,
		UserName: "",
		Mobile:   mobile,
		Email:    "",
		Passwd:   passwd,
	}

	s.CreateUser(&user)

	return nil
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

func (s *UserBasicService) UserSetting(basic model.UserBasic) error {
	// check 该用户是否已经注册过

	userID := basic.UserID
	users := s.GetByUserID(userID)

	if users == nil {
		return errors.New("用户未注册！")
	}

	return nil
}
