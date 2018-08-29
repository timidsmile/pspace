package test

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GoodbyeAction(c *gin.Context) {
	value, exist := c.GetQuery("key")
	if !exist {
		value = "master"
	}
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("goodbye %s !\n", value)))
	return
}
