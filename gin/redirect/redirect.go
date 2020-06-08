package redirect

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Redirect(ctx *gin.Context) {
	//临时重定向
	//ctx.Redirect(http.StatusTemporaryRedirect, "http://baidu.com")
	ctx.Redirect(http.StatusTemporaryRedirect, "/file")
	//ctx.JSON(http.StatusOK, gin.H{
	//	"message": "ok",
}
