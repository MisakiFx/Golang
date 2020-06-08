package firstGin

import "github.com/gin-gonic/gin"

func SayHello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"name": "Misaki",
		"age":  21,
	})
}

//使用RestFul风格Api

func GetBook(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"method": "GET",
		"books":  []string{"C++", "GO", "Java"},
	})
}
func PutBook(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"method":  "PUT",
		"message": "OK",
	})
}
func PostBook(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"method":  "POST",
		"message": "OK",
	})
}
func DeleteBook(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"method":  "DELETE",
		"message": "OK",
	})
}
