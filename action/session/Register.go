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
	response := &components.Response{}
	defer c.JSON(http.StatusOK, response)

	// 支持手机号和邮箱两种注册方式，注册后可以用唯一user_name登陆
	// 直接从表单中取数据形式获取参数
	// email := c.DefaultPostForm("email", "")
	// mobile := c.DefaultPostForm("mobile", "")

	// 把参数取到结构体中，可以指定类型、是否必须
	params := struct {
		Email string `form:"email"`
		Mobile string `form:"mobile"`
		Passwd string `form:"passwd" binding:"required"`
	}{
		// Mobile: "123456", // 赋默认值情形
	}

	if err := c.Bind(&params); err != nil {
		response.Code = 1;
		response.Msg = "邮箱或手机号不能为空!";
		return;
	}

	email := params.Email;
	mobile := params.Mobile;

	ee := components.EmailValidate{}
	mm := components.MobileValidate{}
	if false == ee.Validate(email) || false == mm.Validate(mobile) {
		response.Code = 9999;
		response.Msg = "邮箱或手机号格式不正确!";
		return;
	}

	fmt.Println("email = ", email);
	fmt.Println("mobile = ", mobile);

	passwd := params.Passwd;
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
		*response = userServ.RegisterByEmail(email, passwd, userID);
	} else if(mobile != "") {
		*response = userServ.RegisterByMobile(mobile, passwd, userID);
	} else {
		response.Msg = "仅支持邮箱或手机号注册!";
		return;
	}

	return
}


