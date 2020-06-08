package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP : net.IPv4(127, 0,0,1),
		Port : 9000,
	})
	if err != nil {
		fmt.Println("UdpListen error:", err.Error())
		return
	}
	var data [1024]byte
	var addr *net.UDPAddr
	var n int
	for {
		//接受数据
		n, addr, err = conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read err:", err.Error())
			return
		}
		fmt.Print("clent:", string(data[:n]))
		reply := strings.ToUpper(string(data[:n]))
		//发送数据
		conn.WriteToUDP([]byte(reply), addr)
	}
}
