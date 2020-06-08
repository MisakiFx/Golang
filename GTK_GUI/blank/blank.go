package blank

import "github.com/mattn/go-gtk/gtk"

func Demo() {
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL) //带边框的顶层窗口
	win.SetTitle("go gtk")
	win.SetSizeRequest(480, 320)
	win.Show()
}
