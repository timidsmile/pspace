package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WelcomeAction(c *gin.Context) {
	value, exist := c.GetQuery("key")
	if !exist {
		value = "master!"
	}
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("welcome %s !\n", value)))
	return
}
