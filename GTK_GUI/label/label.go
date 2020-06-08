package label

import (
	"os"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gtk"
)

var (
	State    = true
	Language = "中文"
)

func Label() {
	//创建主窗口
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL) //带边框的顶层窗口
	win.SetTitle("button")
	win.SetSizeRequest(400, 250)
	win.SetTitle("nmsl v2.0")
	win.Connect("destroy", func() {
		gtk.MainQuit()
		os.Exit(0)
	})

	//创建容器控件
	layout := gtk.NewFixed()
	//添加到窗口上
	win.Add(layout)
	//创建标签
	l1 := gtk.NewLabel("nmsl开启")
	l1.SetSizeRequest(200, 100)
	l1.ModifyFontSize(18)
	l1.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("green"))

	l2 := gtk.NewLabel("语言：" + Language)
	l2.SetSizeRequest(200, 100)
	l2.ModifyFontSize(15)
	l2.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("red"))

	l3 := gtk.NewLabel("f8开始互动！")
	l3.SetSizeRequest(400, 50)
	l3.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("purple"))
	l3.ModifyFontSize(15)
	//按钮相关
	b1 := gtk.NewButtonWithLabel("打开/关闭")
	b1.SetSizeRequest(200, 100)
	b1.SetLabelFontSize(10)
	b1.Connect("clicked", func() {
		if State == true {
			State = false
			l1.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("red"))
			l1.SetText("nmsl关闭")
		} else {
			State = true
			l1.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("green"))
			l1.SetText("nmsl开启")
		}
	})

	b2 := gtk.NewButtonWithLabel("切换语言")
	b2.SetSizeRequest(200, 100)
	b2.SetLabelFontSize(10)
	b2.Connect("clicked", func() {
		if Language == "中文" {
			l2.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("blue"))
			Language = "英文"
		} else {
			Language = "中文"
			l2.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("red"))
		}
		l2.SetText("语言：" + Language)
	})

	//组件添加到容器控件中
	layout.Put(l1, 0, 0)
	layout.Put(l2, 0, 100)
	layout.Put(b1, 200, 0)
	layout.Put(b2, 200, 100)
	layout.Put(l3, 0, 200)
	//显示控件
	//layout.Show()
	//也可以
	win.ShowAll()
	//事件处理
	//win.Connect("key-press-event", func(ctx *glib.CallbackContext) {
	//	arg := ctx.Args(0)
	//	event := *(**gdk.EventKey)(unsafe.Pointer(&arg))
	//	key := event.Keyval
	//	fmt.Println(key)
	//})
}
