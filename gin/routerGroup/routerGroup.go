package routerGroup

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouterGroup(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": "non",
	})
}
