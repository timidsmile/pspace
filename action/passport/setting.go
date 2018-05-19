package passport

import (
	"github.com/gin-gonic/gin"
	"github.com/timidsmile/pspace/components"
	"github.com/timidsmile/pspace/model"
	"github.com/timidsmile/pspace/service"
	"net/http"
)

func SettingAction(c *gin.Context) {
	response := components.NewResponse()
	defer c.JSON(http.StatusOK, response)

	// 把参数取到结构体中，可以指定类型、是否必须
	params := struct {
		UserID    int64  `form:"userID" binding:"required"`
		Email     string `form:"email"`
		Mobile    string `form:"mobile"`
		Username  string `form:"userName"`
		NickName  string `form:"nickName"`
		AvatarUrl string `form:"avatarUrl"`
	}{
	// Mobile: "123456", // 赋默认值情形
	}

	// TODO: 如何快速判断输入的所有参数都为空情况
	if err := c.Bind(&params); err != nil {
		response.Code = 1
		response.Msg = "参数不正确!"
		return
	}

	userBasic := model.UserBasic{
		UserID:    params.UserID,
		Email:     params.Email,
		Mobile:    params.Mobile,
		UserName:  params.Username,
		NickName:  params.NickName,
		AvatarUrl: params.AvatarUrl,
	}
	userServ := service.UserBasicService{}
	if err := userServ.UserSetting(userBasic); err != nil {
		response.Code = 1
		response.Msg = err.Error()
	}

	return
}
