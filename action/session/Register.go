package session

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func RegisterAction(c *gin.Context) {
	value, exist := c.GetQuery("key")
	if !exist {
		value = "master"
	}
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("goodbye %s !\n", value)))
	return
}