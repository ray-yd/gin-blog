package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ray-yd/gin-blog/utils/errmsg"
	"io"
	"os"
)

func Upload(c *gin.Context) {
	file, header, _ := c.Request.FormFile("file")
	var code int = errmsg.SUCCESS
	filename := header.Filename
	//fmt.Println(header.Filename)
	out, err := os.Create("./upload/" + filename)
	if err != nil {
		code = errmsg.ERROR
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		code = errmsg.ERROR
	}
	c.JSON(200, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMessage(code),
		"url":     "localhost:3000/" + filename,
	})

}
