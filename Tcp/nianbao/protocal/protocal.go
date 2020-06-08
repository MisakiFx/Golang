package protocal

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

//为了解决tcp粘包，自己写一个协议，前四个字节放数据
func Encode(message string) []byte {
	length := int32(len(message))
	buf := new(bytes.Buffer)
	//这样的方法允许向缓冲区中写入整形
	err := binary.Write(buf, binary.LittleEndian, length)
	if err != nil {
		fmt.Println("18:Encode err:", err.Error())
		return nil
	}
	err = binary.Write(buf, binary.LittleEndian, []byte(message))
	if err != nil {
		fmt.Println("23:Encode err:", err.Error())
		return nil
	}
	return buf.Bytes()
}

func Decode(reader *bufio.Reader) string {
	//这个方法允许只读四个字节
	lengthByte, _:= reader.Peek(4)
	lengthBuffer := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuffer, binary.LittleEndian, &length)
	if err != nil {
		if err == io.EOF {
			return ""
		}
		fmt.Println("Decode error:", err.Error())
		return ""
	}
	if int32(reader.Buffered()) < length {
		fmt.Println("Decode error:", err.Error())
		return ""
	}
	res := make([]byte, length + 4)
	_, err = reader.Read(res)
	if err != nil {
		fmt.Println("Decode error:", err.Error())
		return ""
	}
	return string(res)
}
