package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err:= net.DialUDP("udp", nil, &net.UDPAddr{
		IP : net.IPv4(127,0,0,1),
		Port: 9000,
	})
	if err != nil {
		fmt.Println("connect err:", err.Error())
		return
	}
	defer conn.Close()
	var data [1024]byte
	for {
		reader := bufio.NewReader(os.Stdin)
		line, _:= reader.ReadString('\n')
		conn.Write([]byte(line))
		n, err := conn.Read(data[:])
		if err != nil {
			fmt.Println("read error:", err.Error())
			return
		}
		fmt.Print(string(data[:n]))
	}
}
