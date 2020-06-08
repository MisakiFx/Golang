package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFromFile() {
	//打开文件
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file error : ", err.Error())
		return
	}
	//最后关闭文件
	defer file.Close()
	//读文件
	var cach = make([]byte, 0, 128)
	var b [128]byte
	for {
		count, err := file.Read(b[:])
		if (count <= 0) {
			break
		}
		if err != nil {
			fmt.Println("read file error : ", err.Error())
			break
		}
		cach = append(cach, b[:count]...)
		//string类型只能和切片来回转换，不能转换数组类型
	}
	fmt.Println(string(cach))
}
//使用Buffio这个包来读，可以读取字符串，不用按照byte来读
func readFromBuffio() {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file error : ", err.Error())
		return
	}
	defer file.Close()
	//创建一个用来从文件中读的对象
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("read file error : ", err.Error())
			return
		}
		fmt.Print(line)
	}
}
func readFromFileUiutil() {
	//这个不用打开文件
	//最简单的文件读取方式，直接一次性把文件所有内容读入内存
	ret, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println("read file error:", err.Error())
		return
	}
	fmt.Println(string(ret))
}
//写文件
func writeToFile() {
	file, err := os.OpenFile("./test.txt", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println("open file error:", err.Error())
		return
	}
	defer file.Close()
	file.Write([]byte("\nMisaki牛逼\n"))
	file.WriteString("Misaki牛皮\n")
}
func writeFileBufio() {
	file, err := os.OpenFile("./test.txt", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println("open file error:", err.Error())
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("Misaki牛逼又牛逼")//这个写是写到缓冲区中，不是直接写入文件
	writer.Flush()//刷新缓冲区
}
func WriteFileIoutil() {
	err := ioutil.WriteFile("./test.txt", []byte("Misaki牛皮又牛皮"), 0664)
	if err != nil {
		fmt.Println("write file error:", err.Error())
		return
	}
}
func main() {
	//readFromBuffio()
	//readFromFileUiutil()
	//writeToFile()
	//writeFileBufio()
	WriteFileIoutil()
}
