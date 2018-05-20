package session

import (
	"github.com/gin-gonic/gin"
	"github.com/timidsmile/pspace/components"
	"github.com/timidsmile/pspace/service"
	"net/http"
)

func CheckLoginAction(c *gin.Context) {
	response := components.NewResponse()
	defer c.JSON(http.StatusOK, response)

	token, exist := c.GetPostForm("token")
	if !exist {
		response.Code = 1
		response.Msg = "用户未登录!"
		return
	}

	// 验证token是否正确
	userSession := components.Session{}
	tokenInRedis := userSession.Get(token)

	if tokenInRedis == nil {
		response.Code = 1
		response.Msg = "用户未登录!"
		return
	}

	// 用户已登陆
	userID := tokenInRedis.UserID

	userServ := service.UserBasicService{}
	users := userServ.GetByUserID(userID)

	response.Data = users
	return
}
