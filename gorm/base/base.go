package base

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

//默认用id字段嵌入主键
type User2 struct {
	ID   int64
	Name sql.NullString `gorm:"default:'Misaki'"`
	Age  int64
}
type User struct {
	gorm.Model   //内嵌gorm.Model
	Name         string
	Age          sql.NullInt64 `gorm:"column:user_age"`
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        //设置字段长度255
	MemberNumber *string `gorm:"unique;not null"` //如果设置了默认值，但是传入的为0值则==没传，用指针解决或者用零值
	Num          int     `gorm:"AUTO_INCREMENT"`
	Address      string  `gorm:"index:addr"` //给address设置名为addr的索引
	IgnoreMe     int     `gorm:"-"`          //忽略本字段
}
type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
}

//更改Animal表的名字
func (Animal) TableName() string {
	return "animal"
}
