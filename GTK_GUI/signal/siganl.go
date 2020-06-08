package signal

import (
	"fmt"

	"github.com/mattn/go-gtk/glib"

	"github.com/mattn/go-gtk/gtk"
)

func handleSignal(ctx *glib.CallbackContext) {
	arg := ctx.Data().(string) //获取用户传递的参数
	fmt.Println(arg)
}
func Signal() {
	//创建主窗口
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL) //带边框的顶层窗口
	win.SetTitle("button")
	win.SetSizeRequest(480, 320)

	//创建容器控件
	layout := gtk.NewFixed()
	//添加到窗口上
	win.Add(layout)
	//创建按钮
	b1 := gtk.NewButtonWithLabel("按下去")
	b2 := gtk.NewButtonWithLabel("别按")
	//设置按钮大小
	b2.SetSizeRequest(100, 100)
	//按钮添加到容器控件中
	layout.Put(b1, 0, 0)
	layout.Put(b2, 200, 200)
	//显示控件
	//layout.Show()
	//也可以
	win.ShowAll()
	//创建信号处理
	var str = "Misaki"
	//一点击b1按钮就会触发事件
	b1.Connect("clicked", handleSignal, str)
	//使用命名函数创建信号处理，不用考虑传参
	b2.Connect("clicked", func() {
		fmt.Println(str)
	})
	//窗口关闭事件处理，自动结束进程
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
}
