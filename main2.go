package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args) //环境初始化

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL) //创建窗口
	window.SetPosition(gtk.WIN_POS_CENTER)       //设置窗口居中显示
	window.SetTitle("GTK Go!")                   //设置标题
	window.SetSizeRequest(300, 200)              //设置窗口的宽度和高度

	layout := gtk.NewFixed()

	b1 := gtk.NewButton()
	b1.SetLabel("点击按钮")

	b2 := gtk.NewButtonWithLabel("按钮22222")
	//b2.SetSizeRequest(20, 20)

	window.Add(layout)

	layout.Put(b1, 0, 0)
	layout.Move(b1, 50, 50)
	layout.Put(b2, 50, 100)

	b1.Connect("pressed", func() {
		fmt.Println("我被点击了")
	})

	window.ShowAll()

	gtk.Main() //主事件循环，等待用户操作
}
