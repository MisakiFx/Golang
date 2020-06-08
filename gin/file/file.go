package file

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func File(ctx *gin.Context) {
	//文件上传页面
	ctx.HTML(http.StatusOK, "file.html", nil)
}
func FileUpload(ctx *gin.Context) {
	//从请求中读取文件
	file, err := ctx.FormFile("f1")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "文件有问题",
		})
	}
	//保存本地
	err = ctx.SaveUploadedFile(file, "./"+file.Filename)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "文件保存成功",
	})
}
