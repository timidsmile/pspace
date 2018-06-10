package session

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/timidsmile/pspace/components"
	"github.com/timidsmile/pspace/service"
	"net/http"
	"time"
)

func LoginAction(c *gin.Context) {
	response := components.NewResponse()
	c.Header("Access-Control-Allow-Origin", "http://timidsmile.com")
	defer c.JSON(http.StatusOK, response)

	userName, exist := c.GetPostForm("userName")
	if !exist {
		response.Code = 1
		response.Msg = "请输入用户名!"
		return
	}

	// 使用userName登陆
	userServ := service.UserBasicService{}
	curUser := userServ.GetByUserName(userName)
	if curUser == nil {
		response.Code = 1
		response.Msg = "用户不存在!"
		return
	}

	// 验证密码是否正确
	passwd, exist := c.GetPostForm("passwd")
	if !exist {
		response.Code = 1
		response.Msg = "请输入密码!"
		return
	}

	if passwd != curUser.Passwd {
		response.Code = 1
		response.Msg = "密码不正确!"
		return
	}

	// 登陆成功，记录session
	userID := curUser.UserID
	fmt.Println(userID)

	cur := time.Now()
	timestamp := int64(cur.UnixNano() / 1000000000) //UnitNano获取的是纳秒，除以1000000获取秒级的时间戳
	userSession := components.Session{
		UserID:    userID,
		LoginTime: timestamp,
	}
	token, _ := userSession.Save()
	c.SetCookie("token", token, 86400, "/", "timidsmile.com", false, false)

	fmt.Println("token is :" + token)

	return
}
