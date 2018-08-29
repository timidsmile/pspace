package common

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/timidsmile/pspace/components"
	"github.com/timidsmile/pspace/consts"
)

func UploadAvartarImageAction(c *gin.Context) {
	response := components.NewResponse()
	c.Header("Access-Control-Allow-Origin", "http://www.pspace.com")
	defer c.JSON(http.StatusOK, response)

	fmt.Println("in uploadAvatarImage")
	// single file
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
	}

	// Upload the file to specific dst.
	guid := components.Guid{}
	imgID, _ := guid.NewGUID(consts.DataCenterID_image, consts.WorkID_image)

	idStr := strconv.FormatInt(imgID, 10)

	dst := "./static/images/avatar/" + idStr + ".jpeg"

	fmt.Println(dst)

	url := "/static/images/avatar/" + idStr + ".jpeg"

	if err := c.SaveUploadedFile(file, dst); err != nil {
		fmt.Println(err.Error())
		response.Code = 2
		response.Msg = fmt.Sprintf("'%s' upload failed!", file.Filename)
	} else {
		response.Msg = fmt.Sprintf("'%s' uploaded!", file.Filename)
		response.Data = struct {
			Url string `json:"url"`
		}{Url: url}
	}

	return
}
