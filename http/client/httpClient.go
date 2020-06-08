package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	//response, err := http.Get("http://127.0.0.1:9000/posts/Go/index2?name=Misaki")
	data := url.Values{}
	data.Set("name", "Misaki")
	urlQuery := data.Encode()
	fmt.Println(urlQuery)
	urlObj, _ := url.Parse("http://127.0.0.1:9000/posts/Go/index2")
	urlObj.RawQuery = urlQuery
	req, err := http.NewRequest(http.MethodGet, urlObj.String(), nil)
	//自定义客户端
	//使用短链接
	//client := http.Client{
	//	Transport: &http.Transport{
	//		DisableKeepAlives: true,
	//	},
	//}
	//client.Do(req)
	//默认客户端
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("500...")
		return
	}
	//记着关闭Body
	defer response.Body.Close()
	//var data [1024]byte
	//fmt.Println("status:", response.Status)
	//fmt.Print("Body:")
	//for {
	//	n, _ := response.Body.Read(data[:])
	//	if n == 0  {
	//		break
	//	}
	//	fmt.Print(string(data[:n]))
	//}
	dataRes, err := ioutil.ReadAll(response.Body)
	fmt.Println("Status:", response.Status)
	fmt.Println("Body:", string(dataRes))
}
