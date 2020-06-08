package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	for _, u := range os.Args[1:] {
		urlStuct := url.URL{
			Scheme: "http",
			Host:   u,
		}
		resp, err := http.Get(urlStuct.String())
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		resString := make([]byte, 0, 0)
		for head, value := range resp.Header {
			var headLine = head + ":"
			for _, v := range value {
				headLine += v
				headLine += ";"
			}
			headLine += "\n"
			resString = append(resString, []byte(headLine)...)
		}
		resString = append(resString, []byte("Body:\n")...)
		respBody, err := ioutil.ReadAll(resp.Body)
		resString = append(resString, respBody...)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile("httpBody.txt", resString, 0664)
		if err != nil {
			panic(err)
		}
	}
}
