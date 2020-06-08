package ginJson

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonMap(ctx *gin.Context) {
	//方法1：使用map进行序列化
	//data := map[string]interface{}{
	//	"name":    "Misaki",
	//	"age":     21,
	//	"message": "helloword",
	//}
	ctx.JSON(http.StatusOK, gin.H{
		"name":    "Misaki",
		"age":     21,
		"message": "helloword",
	})
}
func JsonStruct(ctx *gin.Context) {
	type msg struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Message string `json:"message"`
	}
	//data := msg{
	//	Name:    "Misaki",
	//	Age:     21,
	//	Message: "你好",
	//}
	datas := []msg{
		msg{"Misaki", 21, "你好"},
		msg{"Fx", 22, "我很好"},
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": datas,
	})
}
