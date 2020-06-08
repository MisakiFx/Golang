package main

import (
	"bufio"
	"fmt"
	"github.com/MisakiFx/Golang/Tcp/nianbao/protocal"
	"net"
	"sync"
)
//常见的tcp粘包问题，内置缓冲区，数据差不多了一起发，导致无法准确获取客户端单次发送的数据
//解决粘包问题就是自己定一个协议，每次发送在头部给出这次数据的长度
var wg sync.WaitGroup
func process(conn net.Conn) {
	defer wg.Done()
	defer conn.Close()
	reader := bufio.NewReader(conn)
	//temp := [1024]byte{}
	for {
		res := protocal.Decode(reader)
		if res == "" {
			continue
		}
		fmt.Print("一次发送：", res[4:])
	}
}
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("tcpServer:11:", err.Error())
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("tcpServer:11:", err.Error())
			continue
		}
		wg.Add(1)
		go process(conn)
	}
	wg.Wait()
}