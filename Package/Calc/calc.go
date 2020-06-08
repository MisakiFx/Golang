//包名和文件夹名可以不一样
package Calc2

import "fmt"

//这个方法是特殊的方法，一个包中这个方法会在import时自动调用
//无参数无返回值，不能手动调入
//一个包调入时先全局生命该包中的所有全局变量，然后再执行init，如果有main函数最后执行main函数
func init() {
	fmt.Println("import 我自动调入")
}
func Add(num1, num2 int) int {
		return num1 + num2
}