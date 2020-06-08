package main

import (
	"fmt"
	"os"
)

//学生管理系统
//万年不变的练手题材
type student struct {
	id   int64
	name string
}

var (
	allStudent []student = make([]student, 0, 100)
)

func showAllStudent() {
	for _, v := range allStudent {
		fmt.Printf("id:%d, name:%s\n", v.id, v.name)
	}
}
func addStudent() {
	var stu student
	fmt.Printf("请输入学生id:")
	fmt.Scan(&stu.id)
	fmt.Printf("请输入学生name:")
	fmt.Scan(&stu.name)
	allStudent = append(allStudent, stu)
	fmt.Println("添加成功")
}
func deleteStudent() {
	fmt.Printf("请输入删除学生id:")
	var deleteID int64
	fmt.Scan(&deleteID)
	fmt.Printf("要删除的学生id为%v，确认删除么(Y/N)", deleteID)
	var choice string
	fmt.Scan(&choice)
	if choice == "Y" || choice == "y" {
		for i, v := range allStudent {
			if v.id == deleteID {
				allStudent = append(allStudent[:i], allStudent[i+1:]...)
			}
		}
		fmt.Println("删除成功")
	} else {
		fmt.Println("删除取消")
	}
}
func main() {
	for {
		fmt.Printf("欢迎光临学生管理系统\n")
		fmt.Println(`
		1、查看所有学生
		2、新增学生
		3、删除学生
		0、退出
	`)
		fmt.Printf("你要干啥：")
		var choice int
		fmt.Scan(&choice)
		fmt.Println("你选择了", choice)
		switch choice {
		case 0:
			os.Exit(0)
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		default:
			fmt.Println("gun")
		}
	}
}
