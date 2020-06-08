package ginTemplate

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//自定义函数
func Safe(str string) template.HTML {
	return template.HTML(str)
}
func UsersIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "users/index", gin.H{
		"title": "<a href='https://misakifx.github.io'>Misaki的博客</a>",
	})
}
func PostsIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "posts/index", gin.H{
		"title": "posts",
	})
}

func PostsHome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.html", nil)
}
