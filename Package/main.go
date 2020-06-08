package main

//引入永入文件夹，与包名无关
//官方的包在$GOROOT$/src下
//导入从$GOPATH$/src下导入
//匿名包，只想用其中的init方法，而不调用其他接口
import (
	"fmt"
	Calc "github.com/MisakiFx/Golang/Package/Calc"
)
func main() {

	res := Calc.Add(1, 2)
	fmt.Println(res)
}
