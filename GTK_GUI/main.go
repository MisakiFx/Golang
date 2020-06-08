package main

import (
	"bufio"
	"io"
	"os"
	"sync"

	"github.com/MisakiFx/Golang/GTK_GUI/label"
	"github.com/mattn/go-gtk/gtk"

	"github.com/go-vgo/robotgo"
)

var (
	chinese      []string
	english      []string
	chineseIndex = 0
	englishIndex = 0
	wg           sync.WaitGroup
)

//初始化文本
func initText() error {
	chinese = make([]string, 0, 10)
	english = make([]string, 0, 10)
	//初始化中文
	file, err := os.OpenFile("chinese.txt", os.O_RDONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil
		}
		chinese = append(chinese, line)
		if err == io.EOF {
			break
		}
	}
	//初始化英文
	file2, err := os.OpenFile("english.txt", os.O_RDONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer file2.Close()
	reader2 := bufio.NewReader(file2)
	for {
		line, err := reader2.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil
		}
		english = append(english, line)
		if err == io.EOF {
			break
		}
	}
	//for _, str := range chinese {
	//	fmt.Print(str)
	//}
	//for _, str := range english {
	//	fmt.Print(str)
	//}
	return nil
}

//键盘监听
func ListenKey() {
	defer wg.Done()
	for {
		ok := robotgo.AddEvent("f8")
		if ok {
			if label.State == true {
				var text string = ""
				switch label.Language {
				case "中文":
					if len(chinese) <= 0 {
						break
					}
					text = chinese[chineseIndex]
					chineseIndex = (chineseIndex + 1) % len(chinese)
				case "英文":
					if len(english) <= 0 {
						break
					}
					text = (english[englishIndex])
					englishIndex = (englishIndex + 1) % len(english)
				default:

				}
				robotgo.TypeStr(text)
				robotgo.MicroSleep(100)
				robotgo.KeyTap("enter")
			}
		}
		robotgo.MicroSleep(200)
	}
}
func main() {
	err := initText()
	if err != nil {
		return
	}
	wg.Add(1)
	go ListenKey()
	//初始化
	gtk.Init(&os.Args)

	//button.Button()
	//signal.Signal()
	label.Label()
	//主事件循环
	gtk.Main()
	wg.Wait()
}
