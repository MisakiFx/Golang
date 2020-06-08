package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

func main() {
	//创建一个新的缓存
	//设置默认国旗时间为5分钟，每10分钟清除一次过期缓存
	c := cache.New(5*time.Minute, 10*time.Minute)
	c.Set("name", "Misaki", cache.DefaultExpiration)
	v, ok := c.Get("name")
	if ok {
		fmt.Println(v)
	}
}
