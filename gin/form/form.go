package form

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login", nil)
}
func CheckAccount(ctx *gin.Context) {
	//和拿取querystring差不多
	name := ctx.PostForm("username")
	password := ctx.PostForm("password")
	unname := ctx.DefaultPostForm("name", "456")
	temp, ok := ctx.GetPostForm("age")
	if ok {
		name = temp
	}
	ctx.JSON(http.StatusOK, gin.H{
		"username": name,
		"password": password,
		"name":     unname,
	})
}
