package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

//客户端
func main() {
	//与Server建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		if err == io.ErrClosedPipe {
			return
		}
		fmt.Println("client connect err:", err.Error())
		return
	}
	reader := bufio.NewReader(os.Stdin)
	//发送数据
	for {
		line, _ := reader.ReadString('\n')
		conn.Write([]byte(line))
	}
	defer conn.Close()
}
