package queryString

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryString(ctx *gin.Context) {
	//获取浏览器携带的查询字符串
	queryString := ctx.Query("query")
	//能查到就查到，查不到就用默认的
	age := ctx.DefaultQuery("age", "21")
	name, ok := ctx.GetQuery("name")
	if ok {
		queryString = name
	}
	ctx.JSON(http.StatusOK, gin.H{
		"querystring": queryString,
		"age":         age,
	})
}
