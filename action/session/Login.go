package session

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/timidsmile/pspace/service"
	"github.com/timidsmile/pspace/model"
)

func LoginAction(c *gin.Context) {
	value, exist := c.GetQuery("key")
	if !exist {
		value = "master"
	}

	s := service.UserBasicService{}
	ret := s.GetBlogAdmin(1,"123")

	fmt.Println(ret)

	if ret == nil {
		user := &model.UserBasic{
			UserID:"123",
			UserName: "test",
			Mobile: "test",
			Email: "test",
			Passwd: "test",
			NickName: "test",
			AvataUrl: "test",
			Status: 1,
		}

		s.Insert(user);
	}
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("goodbye %s !\n", value)))
	return
}