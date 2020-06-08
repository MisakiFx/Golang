package main

import (
	"github.com/MisakiFx/Golang/gin/file"
	"github.com/MisakiFx/Golang/gin/form"
	"github.com/MisakiFx/Golang/gin/ginJson"
	"github.com/MisakiFx/Golang/gin/ginTemplate"
	"github.com/MisakiFx/Golang/gin/queryString"
	"github.com/MisakiFx/Golang/gin/redirect"
	"github.com/MisakiFx/Golang/gin/routerGroup"
	"github.com/MisakiFx/Golang/gin/shouldBind"
	"github.com/MisakiFx/Golang/gin/uri"
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	//r := gin.Default()
	//建议使用Restful风格API
	//r.GET("/hello", firstGin.SayHello)
	//r.GET("/book", firstGin.GetBook)
	//r.PUT("/book", firstGin.PutBook)
	//r.POST("/book", firstGin.PostBook)
	//r.DELETE("/book", firstGin.DeleteBook)
	//r.Run(":9000")
	//r := gin.Default()
	//r.Use()
	//r.GET()
	//http.HandleFunc("/template", template.SayHello)
	//http.HandleFunc("/template2", template.SayHello2)
	//http.HandleFunc("/template3", template.SayHello3)
	//http.HandleFunc("/template4", template.SayHello4)
	//http.HandleFunc("/selfFunc", template.SelfFunc)
	//http.HandleFunc("/qiantao", template.QianTao)
	//http.HandleFunc("/index", template.Index)
	//http.HandleFunc("/home", template.Home)
	//err := http.ListenAndServe(":9000", nil)
	//if err != nil {
	//	fmt.Println("server error:", err.Error())
	//	return
	//}
	r := gin.Default()
	//与http都在handler中解析并渲染不同
	//在外部解析，只解析一遍
	//r.LoadHTMLFiles("ginTemplate/index.tmpl")
	//r.LoadHTMLFiles("ginTemplate/posts/index.tmpl", "ginTemplate/users/index.tmpl")
	//给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": ginTemplate.Safe,
	})
	//解析文件
	//r.LoadHTMLFiles("form/login.html")
	r.LoadHTMLGlob("ginTemplate/**/*")
	//加载静态文件
	r.Static("/static", "./static")
	//在handler里面渲染
	r.GET("/posts/index", ginTemplate.PostsIndex)
	r.GET("/users/index", ginTemplate.UsersIndex)
	//返回前端模板网页
	r.GET("/home", ginTemplate.PostsHome)
	//返回json格式
	r.GET("/jsonMap", ginJson.JsonMap)
	r.GET("/jsonStruct", ginJson.JsonStruct)
	//querystring
	r.GET("/queryString", queryString.QueryString)
	//form表单获取
	r.GET("/login", form.Login)
	r.POST("/login/check", form.CheckAccount)
	//uri参数获取
	r.GET("/uri/:name/:age", uri.Uri)
	//shouldBind自动绑定结构体
	r.GET("/shouldBind", shouldBind.ShouldBind)
	r.POST("/shouldBindPost", shouldBind.ShouldBindPost)
	//文件上传
	r.GET("/file", file.File)
	r.POST("/file", file.FileUpload)
	//请求重定向
	r.GET("/redirect", redirect.Redirect)
	//路由与路由组
	r.GET("/routerGroup", routerGroup.RouterGroup)
	//事先没有准备好的地址和方法，会调用这个句柄
	r.NoRoute(routerGroup.RouterGroup)
	//路由组，公用的前缀，提取出来成为一个路由组
	videoGroup := r.Group("/video")
	{
		//加个代码块，不是必须的，但是看起来很整齐，很舒服
		videoGroup.GET("/index", file.File)
		videoGroup.GET("/home", shouldBind.ShouldBind)
	}
	//路由组下面还可以有路由组。。等等
	r.Run(":9000")
}
