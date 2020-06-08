package main

import "fmt"

func main() {
	//复杂数据类型
	//数组，连续的
	//var b [4]int
	//初始化方式1
	//var a [3]int = [3]int{1, 1, 1} //不赋值都是0值
	//b = [4]int{1, 1, 1}
	//初始化方式2，自动推断初始化数据长度给空间
	c := [...]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(c); i++ {
		fmt.Printf("%d ", c[i])
	}
	fmt.Println()
	//初始化方式3，根据索引初始化，不初始化的部分为0
	d := [5]int{0: 1, 4: 2}
	for _, v := range d {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	//多维数组
	var aa [3][3]int
	aa = [3][3]int{
		[3]int{1, 2, 3},
		[3]int{3, 4, 5},
		[3]int{4, 5, 6},
	}
	for _, v := range aa {
		fmt.Printf("%d ", v)
		fmt.Println()
	}
	for i := 0; i < len(aa); i++ {
		for j := 0; j < len(aa[i]); j++ {
			fmt.Printf("%d ", aa[i][j])
		}
		fmt.Println()
	}
	//在go中数组是基本类型，所以允许直接使用==和!=比较以及数组变量之间的拷贝
	//就是方便啊
	//切片是一个引用，是底层数组的引用，也是连续的
	//切片的底层就是一个数组，无论是从已有数组上切还是新建切片
	//都依赖于一个数组
	var s1 []int //不写长度就是切片了
	var s2 []string
	s1 = []int{1, 2, 3}
	s2 = []string{"Misaki", "Fx"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) //nil在go终究是没有开辟内存空间就是类似于C++中的null空指针
	fmt.Printf("len s1 = %d, cap s1 = %d\n", len(s1), cap(s1))
	//长度
	fmt.Printf("s1的长度为%d, s2的长度为%d\n", len(s1), len(s2))
	//由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9}
	s3 := a1[1:4] //左闭右开区间
	fmt.Println(s3)
	s4 := a1[:4] //从零开始切到第四个
	//越界
	//s4[4] = 5
	//hfmt.Printf("a1 = %v\n", a1)
	var s5 []int = a1[1:] //从1开始切到最后一个
	//s6 := a1[:]  //梭哈
	//cap为取容量，切片的容量为底层数组从切片第一个元素到最后一个元素的数量
	fmt.Printf("len s4 = %d, cap s4 = %d\n", len(s4), cap(s4))
	fmt.Printf("len s5 = %d, cap s5 = %d\n", len(s5), cap(s5))
	//切片再切片
	//长度和容量只和原有数组有关
	s7 := s4[1:]
	fmt.Printf("s7 = %v, len s7 = %d, cap s7 = %d\n", s7, len(s7), cap(s7))
	//原有数组改变切片都会改变，反之一样
	a1[3] = 100
	fmt.Printf("s7 = %v, len s7 = %d, cap s7 = %d\n", s7, len(s7), cap(s7))
	s7[2] = 1
	fmt.Printf("a1 = %v, s4 = %v\n", a1, s4)
	//用make创建切片
	ms1 := make([]int, 5, 10)
	fmt.Printf("%v: len = %d, cap = %d\n", ms1, len(ms1), cap(ms1))
	//切片不能直接比较，不能用==操作符来判断两个切片中的元素是否相同，唯一合法比较是和nil比较
	//nil切片长度为0，但长度为0的切片不一定是nil比如一下的例子
	var emptys1 []int = []int{}
	emptys2 := make([]int, 0)
	var emptys3 []int
	fmt.Printf("%v:len = %d, cap = %d, emptys1 == nil = %v\n", emptys1, len(emptys1), cap(emptys1), emptys1 == nil)
	fmt.Printf("%v:len = %d, cap = %d, emptys2 == nil = %v\n", emptys2, len(emptys2), cap(emptys2), emptys2 == nil)
	fmt.Printf("%v:len = %d, cap = %d, emptys3 == nil = %v\n", emptys3, len(emptys3), cap(emptys3), emptys3 == nil)
	//为切片追加元素
	ss1 := []string{"123", "456", "789"}
	fmt.Printf("%v, len = %d, cap = %d\n", ss1, len(ss1), cap(ss1))
	//用append必须用原来的切片接收它
	ss1 = append(ss1, "11123")
	//可以发现cap容量也自动扩大了，说明底层数组扩容了，原理类似vector的push_back
	//扩容策略，新容量大于原来容量二倍就是新容量
	//小于二倍如果原来的长度小于1024则容量就是旧容量的二倍
	//如果大于1024则新容量每次增加原来旧容量的1/4
	//如果最终容量计算溢出，则用新申请的容量
	fmt.Printf("%v, len = %d, cap = %d\n", ss1, len(ss1), cap(ss1))
	ss1 = append(ss1, "成都", "杭州")
	fmt.Printf("%v, len = %d, cap = %d\n", ss1, len(ss1), cap(ss1))
	ss2 := []string{"Misaki", "张三", "李四"}
	ss1 = append(ss1, ss2...) //...表示拆开
	fmt.Printf("%v, len = %d, cap = %d\n", ss1, len(ss1), cap(ss1))
	//copy()复制切片
	sss1 := []int{1, 2, 3}
	sss2 := sss1
	sss3 := make([]int, 3, 3)
	//拷贝切片是将一个切片中的值拷贝进另一个切片，这两个切片底层数组可以并不是同一个数组，存在于不同内存空间
	copy(sss3, sss1)
	fmt.Printf("%v, %v, %v\n", sss1, sss2, sss3)
	sss1[0] = 0
	fmt.Printf("%v, %v, %v\n", sss1, sss2, sss3)
	//从切片中删除元素
	//切片本身并没有提供删除元素的接口，但是我们可以借助append来删除元素
	//删除下表为2的元素
	sss4 := []int{1, 2, 3, 4}
	fmt.Println(sss4)
	sss4 = append(sss4[:2], sss4[3:]...)
	fmt.Println(sss4)
	//这一套操作过后底层数组的容量不发生变化，也没有内存的申请和释放，是优秀的删除方法
	fmt.Printf("len(sss4) = %d, cap(sss4) = %d\n", len(sss4), cap(sss4))
	//直观的验证
	aa1 := [...]int{1, 2, 3}
	sss5 := aa1[:]
	fmt.Printf("aa1 = %v, len = %d, cap = %d\n", aa1, len(aa1), cap(aa1))
	fmt.Printf("sss5 = %v, len = %d, cap = %d\n", sss5, len(sss5), cap(sss5))
	sss5 = append(sss5[:1], sss5[2:]...)
	//可见底层数组的变化，十分巧妙
	fmt.Printf("aa1 = %v, len = %d, cap = %d\n", aa1, len(aa1), cap(aa1))
	fmt.Printf("sss5 = %v, len = %d, cap = %d\n", sss5, len(sss5), cap(sss5))
}
