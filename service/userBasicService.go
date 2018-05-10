package service

import (
	"github.com/timidsmile/pspace/components"
	"sync"
	"github.com/timidsmile/pspace/model"
)

type UserBasicService struct {
	mutex *sync.Mutex
}

func (u *UserBasicService) Insert(user *model.UserBasic) error {
	tx := components.Db.Begin()
	if err := tx.Create(user).Error; nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}


func (srv *UserBasicService) GetBlogAdmin(id uint64, userID string) *model.UserBasic {
	ret := &model.UserBasic{}
	if err := components.Db.Where("`id` = ? AND `user_id` = ?",
		1,userID).Order("`id` asc").First(ret).Error; nil != err {

		return nil
	}

	return ret
}