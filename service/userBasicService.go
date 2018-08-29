package service

import (
	"github.com/timidsmile/pspace/components"
	"github.com/timidsmile/pspace/model"
	"sync"
)

type userBasicService struct {
	mutex *sync.Mutex
}

func (u *userBasicService) insert(user *model.User) error {
	tx := components.Db.Begin()
	if err := tx.Create(user).Error; nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}
