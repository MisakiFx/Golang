package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)
var wg sync.WaitGroup
func process(conn net.Conn) {
	//与客户端通讯
	defer wg.Done()
	temp := [128]byte{}
	for {
		n, err := conn.Read(temp[:])
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("read err:", err.Error())
			conn.Close()
			return
		}
		fmt.Print(string(temp[:n]))
	}
}
//tcp服务端
func main() {
	//1、本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("Listen error:", err.Error())
		return
	}
	//2、等待别人来跟我建立连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("connect err:", err.Error())
			return
		}
		//用一个groutine去处理每个连接
		wg.Add(1)
		go process(conn)
	}
	wg.Wait()
}
