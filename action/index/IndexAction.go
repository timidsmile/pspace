package index

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func IndexAction(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.HTML(200, "login.html", pongo2.Context{})
}
