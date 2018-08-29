package passport

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/timidsmile/pspace/components"
	"github.com/timidsmile/pspace/service"
)

func GetUserInfoAction(c *gin.Context) {
	response := components.NewResponse()
	c.Header("Access-Control-Allow-Origin", "http://www.pspace.com")
	defer c.JSON(http.StatusOK, response)

	// 直接从表单中取数据形式获取参数
	userID := c.MustGet("userID").(int64)
	fmt.Println(userID)

	userServ := service.UserBasicService{}
	userBasic := userServ.GetByUserID(userID)

	response.Data = userBasic
	return
}
