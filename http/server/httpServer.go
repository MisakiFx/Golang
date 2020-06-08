package server

//go内置的http库十分强大
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(response http.ResponseWriter, request *http.Request) {
	//str := `<h1 style="color:red">Misaki<h1>`

	str, _ := ioutil.ReadFile("./xx.html")
	response.Write(str)
}
func f2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println(r.URL)
	//fmt.Println(r.Header.Get("Content-Type"))
	fmt.Println(r.URL.Query())
	fmt.Println(r.URL.Query().Get("name"))
	w.Write([]byte("Misaki"))
}
func main() {
	http.HandleFunc("/posts/Go/index", f1)
	http.HandleFunc("/posts/Go/index2", f2)
	http.ListenAndServe("127.0.0.1:9000", nil)
}
