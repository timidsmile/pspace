package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/timidsmile/pspace/components"
	"net/http"
)

func UploadFileAction(c *gin.Context) {
	response := components.NewResponse()
	c.Header("Access-Control-Allow-Origin", "http://www.pspace.com")
	defer c.JSON(http.StatusOK, response)

	fmt.Println("in")
	// single file
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
	}

	// Upload the file to specific dst.
	if err := c.SaveUploadedFile(file, "./static/images/2.jpeg"); err != nil {
		fmt.Println(err.Error())
	}

	response.Msg = fmt.Sprintf("'%s' uploaded!", file.Filename)

	return
}
