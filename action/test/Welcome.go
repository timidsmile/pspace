package test

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomeAction(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	value, exist := c.GetQuery("key")
	if !exist {
		value = "master!"
	}
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("welcome %s !\n", value)))
	return
}
