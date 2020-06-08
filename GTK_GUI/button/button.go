package button

import "github.com/mattn/go-gtk/gtk"

func Button() {
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
}
