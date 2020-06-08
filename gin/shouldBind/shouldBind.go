package shouldBind

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type user struct {
	Name string `form:"name" json:"name"`
	Age  int    `form:"age" json:"age1"`
}

func ShouldBind(ctx *gin.Context) {
	var u user
	//必须传指针
	err := ctx.ShouldBind(&u)
	fmt.Println(u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ok",
		})
	}
}
func ShouldBindPost(ctx *gin.Context) {
	var u user
	//必须传指针
	err := ctx.ShouldBind(&u)
	fmt.Println(u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}
