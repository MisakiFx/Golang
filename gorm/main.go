package main

import (
	"fmt"
	"github.com/MisakiFx/Golang/gorm/base"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//数据表
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	//连接数据库
	db, err := gorm.Open("mysql", "root:1@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&base.User2{})
	u := base.User2{
		Age: 21,
	}
	fmt.Println(db.NewRecord(&u)) //判断是否是一条新数据
	db.Create(&u)
	fmt.Println(db.NewRecord(&u)) //判断是否是一条新数据
	db.Create(&u)
	////禁用复数形式
	//db.SingularTable(true)
	////创建表 自动迁移(自动和数据库中的表对应)
	//db.AutoMigrate(&base.User{})
	//db.AutoMigrate(&UserInfo{})
	//db.AutoMigrate(&base.Animal{})
	////用名字创建一个表
	//db.Table("Misaki").CreateTable(&UserInfo{})

	//插入数据
	//创建数据行
	//u1 := UserInfo{
	//	1,
	//	"Misaki",
	//	"男",
	//	"蛙泳",
	//}
	//db.Create(&u1)

	//查询
	//var u UserInfo
	//db.First(&u)
	//fmt.Println(u)

	//更新
	//db.Model(&u).Update("hobby", "游戏")

	//删除
	//db.Delete(&u)
}
