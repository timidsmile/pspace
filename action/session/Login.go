package session

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginAction(c *gin.Context) {
	value, exist := c.GetQuery("key")
	if !exist {
		value = "master"
	}
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("goodbye %s !\n", value)))
	return
}
