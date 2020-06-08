package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "src\\Misaki"
	fmt.Println(path)
	//多行字符串，``中的内容原样输出
	s2 := `
		Misaki
		Misaki2
		Misaki3
	`
	fmt.Println(s2)
	//字符串常用操作
	//长度
	fmt.Println(len(path))
	s3 := "Misaki2"
	//拼接
	ss := s3 + path                      //+号拼接
	ss2 := fmt.Sprintf("%s%s", s3, path) //fmt拼接，与Printf不同的是不打印二十返回一个字符串
	fmt.Println(ss)
	fmt.Println(ss2)
	//分割
	ret := strings.Split(ss, "\\")
	fmt.Println(ret)
	//包含
	fmt.Println(strings.Contains(ss, "Misaki"))
	//前缀
	fmt.Println(strings.HasPrefix(ss, "Misaki"))
	//后缀
	fmt.Println(strings.HasSuffix(ss, "Misaki"))
	//字串位置
	fmt.Println(strings.Index(ss, "Misaki"))
	fmt.Println(strings.LastIndex(ss, "Misaki"))
	//拼接切片
	fmt.Println(strings.Join(ret, "+"))
	//byte和rune类型，go中没有char类型，用byte和rune分别表示中英文字符
	//字符串无法进行修改，要想修改必须变成byte类型或者rune类型，然后再转换为string进行使用
	//byte类型代指英文字符，rune指其它字符
	//byte是int8，rune是int32
	s1 := "big"
	byte1 := []rune(s1)
	byte1[0] = 'p'
	fmt.Printf("%s\n", string(byte1))
	s9 := "张三"
	rune1 := []rune(s9)
	rune1[0] = '李'
	fmt.Printf("%s\n", string(rune1))
	//%c占位符可以用于代表byte和rune两种类型
	c1 := rune('张')
	c2 := byte('M')
	fmt.Printf("%c\n%c\n", c1, c2)
}
