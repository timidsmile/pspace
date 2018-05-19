package session

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/timidsmile/pspace/components"
	"github.com/timidsmile/pspace/consts"
	"github.com/timidsmile/pspace/service"
	"net/http"
)

func RegisterAction(c *gin.Context) {
	response := components.NewResponse()
	defer c.JSON(http.StatusOK, response)

	// 支持手机号和邮箱两种注册方式，注册后可以用唯一user_name登陆
	// 直接从表单中取数据形式获取参数
	email := c.DefaultPostForm("email", "")
	mobile := c.DefaultPostForm("mobile", "")
	passwd := c.PostForm("passwd")

	if passwd == "" {
		response.Code = 1
		response.Msg = "请先设置密码"
		return
	}

	ee := components.EmailValidate{}
	mm := components.MobileValidate{}
	if false == ee.Validate(email) && false == mm.Validate(mobile) {
		response.Code = 9999
		response.Msg = "邮箱或手机号格式不正确!"
		return
	}

	fmt.Println("email = ", email)
	fmt.Println("mobile = ", mobile)

	// 校验密码格式

	// 生成用户id
	guid := components.Guid{}
	userID, _ := guid.NewGUID(consts.DataCenterID_session, consts.WorkID_session)

	userServ := service.UserBasicService{}

	var err error
	if email != "" {
		err = userServ.RegisterByEmail(email, passwd, userID)
	} else if mobile != "" {
		err = userServ.RegisterByMobile(mobile, passwd, userID)
	} else {
		err = errors.New("仅支持邮箱或手机号注册")
	}

	if err != nil {
		response.Code = 1
		response.Msg = err.Error()
	}
	return
}
