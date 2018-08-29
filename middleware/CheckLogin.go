package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/timidsmile/pspace/components"
)

func CheckLogin(c *gin.Context) {

	if cookie, err := c.Request.Cookie("token"); err == nil {
		token := cookie.Value

		if token == "" {
			response := components.NewResponse()
			response.Code = 1
			response.Msg = "用户未登录!"
			c.JSON(http.StatusOK, response)

			c.Abort()
			return
		}

		userSession := components.Session{}
		tokenInRedis := userSession.Get(token)
		if tokenInRedis == nil {
			response := components.NewResponse()
			response.Code = 1
			response.Msg = "用户未登录!"
			c.JSON(http.StatusOK, response)

			c.Abort()
			return
		}

		userID := tokenInRedis.UserID
		fmt.Println(userID)
		c.Set("userID", userID)
	}

	c.Next()

	return
}
