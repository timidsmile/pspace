package passport

import (
	"github.com/gin-gonic/gin"
	"github.com/timidsmile/pspace/components"
	"net/http"
	"github.com/timidsmile/pspace/service"
	"strconv"
	"fmt"
)

func GetUserInfoAction(c *gin.Context) {
	fmt.Println("in getU")
	response := components.NewResponse()
	c.Header("Access-Control-Allow-Origin", "*")
	defer c.JSON(http.StatusOK, response)

	// 直接从表单中取数据形式获取参数
	userIdStr := c.PostForm("userID")

	userID, _ := strconv.ParseInt(userIdStr, 10, 64)

	userServ := service.UserBasicService{}
	userBasic := userServ.GetByUserID(userID)

	response.Data = userBasic;
	return
}
