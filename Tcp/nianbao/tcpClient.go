package main

import (
	"fmt"
	"github.com/MisakiFx/Golang/Tcp/nianbao/protocal"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	defer conn.Close()
	if err != nil {
		fmt.Println("connect error:", err.Error())
		return
	}
	for i := 0; i < 20; i++ {
		//conn.Write([]byte("Misaki\n"))
		//用协议解决粘包
		conn.Write(protocal.Encode("Misaki"))
	}
}
