package session

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/timidsmile/pspace/service"
	"github.com/timidsmile/pspace/components"
	"github.com/timidsmile/pspace/consts"
)

func RegisterAction(c *gin.Context) {
	response := components.Response{}
	defer c.JSON(http.StatusOK, response)

	// 支持手机号和邮箱两种注册方式，注册后可以用唯一user_name登陆
	email := c.DefaultPostForm("email", "")
	mobile := c.DefaultPostForm("mobile", "")


	if email == "" && mobile == "" {
		response.Code = 1;
		response.Msg = "邮箱或手机号不能为空";
		return;
	}

	// TODO:验证邮箱格式或手机号格式

	fmt.Println("email = ", email);
	fmt.Println("mobile = ", mobile);

	passwd, _:= c.GetPostForm("passwd")
	if passwd == "" {
		response.Code = 1;
		response.Msg = "请先设置密码";
		return;
	}

	// 校验密码格式


	// 生成用户id
	guid := components.Guid{};
	userID,_ := guid.NewGUID(consts.DataCenterID_session, consts.WorkID_session);

	userServ := service.UserBasicService{};

	if(email != "") {
		response = userServ.RegisterByEmail(email, passwd, userID);
	} else if(mobile != "") {
		response = userServ.RegisterByMobile(mobile, passwd, userID);
	} else {
		response.Msg = "仅支持邮箱或手机号注册!";
		return;
	}

	return
}


