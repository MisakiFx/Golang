package main

import (
	"encoding/json"
	"fmt"
	"time"
)

//序列化和反序列化
//这里要大写，不然对外部包不可见，相当于private
//或者使用json特殊解析格式
type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}
type guiyin struct {
	Name []string `json:"name"`
}
type whiteList struct {
	List map[int64]map[string][]string `json:"list"`
}
type test struct {
	CreateTime time.Time
	Id         *int
}
type Base struct {
	Code    int
	Message string
}
type Resp struct {
	BaseResp Base
	Count    int
	Configs  []string
}

func main() {
	//list := make(map[int]*string)
	//s := "123"
	//list[0] = &s
	//jsonString, _ := json.Marshal(list)
	//fmt.Println(string(jsonString))
	resp := Resp{
		BaseResp: Base{
			Code:    0,
			Message: "null",
		},
		Count:   10,
		Configs: []string{"config1", "config2"},
	}
	jsonString, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonString))
	//list := make(map[int64]map[string][]string)
	//var list2 map[int64]map[string][]string
	//domain := make(map[string][]string)
	//domain["www.toutiao.com"] = []string{
	//	"__MAC__",
	//	"__IMEI__",
	//}
	//domain["www.tiktok.com"] = []string{
	//	"__OS__",
	//	"__AG__",
	//}
	//list[123456] = domain
	//list[789] = domain
	////white := whiteList{list}
	//jsonString, err := json.Marshal(list)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(jsonString))
	//err = json.Unmarshal(jsonString, &list2)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(list2)
	//fmt.Println(list2[123456]["www.tiktok.com"])
	//list := make(map[string]map[int]string)
	//list2 := make(map[string]map[int]string)
	//xiansuolei := make(map[int]string)
	//xiansuolei[0] = "表单提交"
	//xiansuolei[1] = "有效咨询"
	//xiansuolei[2] = "留资咨询"
	//douyinhao := make(map[int]string)
	//douyinhao[0] = "私信"
	//douyinhao[1] = "私信点击"
	//douyinhao[2] = "购物车"
	//list["线索类"] = xiansuolei
	//list["抖音号推广"] = douyinhao
	//jsonString, err := json.Marshal(list)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(jsonString))
	//err = json.Unmarshal(jsonString, &list2)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(list2)
	//var test1 test
	//var test2 test
	//test1.Id = 1
	//jsonString, _ := json.Marshal(test1)
	//fmt.Println(string(jsonString))
	//test1.CreateTime, _ = time.Parse("2006-01-02 15:04:05", "1999-01-20 20:00:00")
	//fmt.Println(test1.CreateTime)
	//jsonString, _ := json.Marshal(test1)
	//fmt.Println(string(jsonString))
	//json.Unmarshal(jsonString, &test2)
	//fmt.Println(test2.CreateTime)
	//err := json.Unmarshal([]byte(`{"id":"1"}`), &test1)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(test1.Id)
}

//func main() {
//	p1 := person {
//		Name : "Misaki",
//		Age : 21,
//	}
//	jsonString, err := json.Marshal(p1)
//	if err != nil {
//		fmt.Println("error", err.Error())
//		return
//	}
//	fmt.Printf("%s\n", string(jsonString))
//	str := `{"name":"Misaki","age":21}`
//	var p2 person
//	//反序列化
//	json.Unmarshal([]byte(str), &p2)
//	fmt.Println(p2)
//	jsonString = []byte(`{"name":["Misaki", "MisakiFx"]}`)
//	var g guiyin
//	json.Unmarshal(jsonString, &g)
//	fmt.Println(g)
//}
