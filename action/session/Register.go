package session

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/timidsmile/pspace/components"
	"github.com/timidsmile/pspace/consts"
	"github.com/timidsmile/pspace/model"
	"github.com/timidsmile/pspace/service"
)

func RegisterAction(c *gin.Context) {
	response := components.NewResponse()
	c.Header("Access-Control-Allow-Origin", "*")
	defer c.JSON(http.StatusOK, response)

	// 使用用户名注册方式
	// 至少绑定手机号和邮箱之一，方便以后找回密码使用。
	// 注册后可以用唯一user_name登陆

	// 直接从表单中取数据形式获取参数
	//email := c.DefaultPostForm("userName", "")
	//passwd := c.PostForm("passwd")

	params := struct {
		UserName string `form:"userName" binding:"required" `
		Email    string `form:"email"`
		Mobile   string `form:"mobile"`
		Passwd   string `form:"passwd" binding:"required"`
	}{
		Mobile: "", // 赋默认值情形
		Email:  "",
	}

	if err := c.Bind(&params); err != nil {
		response.Code = 1
		response.Msg = "参数不正确!"
		return
	}

	userName := params.UserName
	// userName 校验格式正确性
	userNameReg := components.UserNameValidate{}
	if false == userNameReg.Validate(userName) {
		response.Code = 9999
		response.Msg = "用户名仅支持字母[a-zA-z]、数字[0-9]、下划线[_]，最少4位且不得超过32位!"
		return
	}

	email := params.Email
	mobile := params.Mobile

	emailReg := components.EmailValidate{}
	mobileReg := components.MobileValidate{}
	if false == emailReg.Validate(email) && false == mobileReg.Validate(mobile) {
		response.Code = 9999
		response.Msg = "邮箱或手机号格式不正确!"
		return
	}

	fmt.Println("email = ", email)
	fmt.Println("mobile = ", mobile)

	passwd := params.Passwd
	// 校验密码格式
	if passwd == "" {
		response.Code = 1
		response.Msg = "请先设置密码"
		return
	}

	// 生成用户id
	guid := components.Guid{}
	userID, _ := guid.NewGUID(consts.DataCenterID_session, consts.WorkID_session)

	userBasic := model.UserBasic{
		UserID:    userID,
		UserName:  userName,
		Mobile:    mobile,
		Email:     email,
		Passwd:    passwd,
		NickName:  "",
		AvatarUrl: "",
		Status:    1,
	}

	userServ := service.UserBasicService{}

	if err := userServ.Register(userBasic); err != nil {
		response.Code = 1
		response.Msg = err.Error()
	}

	// set session
	cur := time.Now()
	timestamp := int64(cur.UnixNano() / 1000000000) //UnitNano获取的是纳秒，除以1000000获取秒级的时间戳
	userSession := components.Session{
		UserID:    userID,
		LoginTime: timestamp,
	}

	token, _ := userSession.Save()

	c.SetCookie("token", token, 1111111111111111, "/", "www.pspace.com", true, true)

	c.String(http.StatusOK, "register successful")

	return
}
