package uri

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Uri(ctx *gin.Context) {
	name := ctx.Param("name")
	age := ctx.Param("age")
	ctx.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}
