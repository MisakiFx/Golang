package template

import (
	"fmt"
	"html/template"
	"net/http"
)

//操作模板
func SayHello(w http.ResponseWriter, r *http.Request) {
	//1、定义模板，tmpl文件
	//2、解析模板，
	t, err := template.ParseFiles("template/hello.tmpl")
	if err != nil {
		fmt.Println("parse error:", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("<h1>解析错误</h1>"))
		return
	}
	//3、设置模板
	err = t.Execute(w, "Misaki")
	if err != nil {
		fmt.Println("Set error:", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("<h3>设置错误</h3>"))
		return
	}
}

type user struct {
	Name   string
	Gender string
	Age    int
}

//传入一个结构体
func SayHello2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/hello2.tmpl")
	if err != nil {
		fmt.Println("parse error:", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("<h1>解析错误</h1>"))
		return
	}
	err = t.Execute(w, user{
		"Misaki",
		"男", //小写在模板中不能访问
		21,
	})
	if err != nil {
		fmt.Println("Set error:", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("设置错误"))
		return
	}
}

//向模板中传入一个map
func SayHello3(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/hello3.tmpl")
	if err != nil {
		fmt.Println("parse error:", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("<h3>解析错误</h3>"))
		return
	}
	m1 := map[string]interface{}{
		"name":   "Misaki",
		"gender": "男",
		"age":    21,
	}
	err = t.Execute(w, m1)
	if err != nil {
		fmt.Println("Set error:", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("设置错误"))
		return
	}
}

//合并使用，传入多个值
//推荐使用map合并，更加方便
func SayHello4(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/hello4.tmpl")
	if err != nil {
		fmt.Println("parse error:", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("<h3>解析错误</h3>"))
		return
	}
	//向模板中传入一个map
	m1 := map[string]interface{}{
		"name":   "Misaki",
		"gender": "男",
		"age":    21,
	}
	u1 := user{
		"Misaki",
		"男", //小写在模板中不能访问
		21,
	}
	err = t.Execute(w, map[string]interface{}{
		"u1": u1,
		"m1": m1,
	})
	if err != nil {
		fmt.Println("Set error:", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("设置错误"))
		return
	}
}
func SelfFunc(w http.ResponseWriter, r *http.Request) {
	//1、定义模板，tmpl文件
	//我们自定义一个函数
	//要么一个返回值要么两个返回值，第二个必须是error类型
	kua := func(name string) (string, error) {
		return name + "牛逼", nil
	}
	//2、解析模板，使用一个新的方法解析
	t := template.New("selfFunc.tmpl") //这里只能写个文件名，不能写路径
	//自定义函数要在解析模板之前
	t.Funcs(template.FuncMap{
		"kua": kua, //前面是模板中可以用的函数名，后面是实际的函数
	})
	t, err := t.ParseFiles("template/selfFunc.tmpl")
	if err != nil {
		fmt.Println("parse error:", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("<h1>解析错误</h1>"))
		return
	}
	//3、设置模板
	err = t.Execute(w, "Misaki")
	if err != nil {
		fmt.Println("Set error:", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("<h1>设置错误</h1>"))
		return
	}
}
func QianTao(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/qiantao.tmpl", "template/hello.tmpl")
	if err != nil {
		fmt.Println("parse error:", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("<h1>解析错误</h1>"))
		return
	}
	err = t.Execute(w, "Misaki")
	if err != nil {
		fmt.Println("Set error:", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("<h1>设置错误</h1>"))
		return
	}
}
func Index(w http.ResponseWriter, r *http.Request) {
	//使用模板就得解析两个模板
	t, err := template.ParseFiles("template/templateDeliver/base.tmpl", "template/templateDeliver/index.tmpl")
	if err != nil {
		fmt.Println("parse error:", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("<h1>解析错误</h1>"))
		return
	}
	err = t.ExecuteTemplate(w, "index.tmpl", "index")
	if err != nil {
		fmt.Println("Set error:", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("<h1>设置错误</h1>"))
		return
	}
}
func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/templateDeliver/base.tmpl", "template/templateDeliver/home.tmpl")
	if err != nil {
		fmt.Println("parse error:", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("<h1>解析错误</h1>"))
		return
	}
	err = t.ExecuteTemplate(w, "home/tmpl", "home")
	if err != nil {
		fmt.Println("Set error:", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("<h1>设置错误</h1>"))
		return
	}
}
