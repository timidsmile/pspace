package session

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/timidsmile/pspace/model"
	"github.com/timidsmile/pspace/service"
	"net/http"
)

func LoginAction(c *gin.Context) {
	value, exist := c.GetQuery("key")
	if !exist {
		value = "master"
	}

	s := service.UserBasicService{}
	ret := s.GetByUserID(123)

	fmt.Println(ret)

	if ret == nil {
		user := &model.UserBasic{
			UserID:    123,
			UserName:  "test",
			Mobile:    "test",
			Email:     "test",
			Passwd:    "test",
			NickName:  "test",
			AvatarUrl: "test",
			Status:    1,
		}

		s.CreateUser(user)

	}
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("goodbye %s !\n", value)))
	return
}
