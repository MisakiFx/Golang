package server

import (
	"fmt"
	"net/http"
	"sync"
)

var count = 0
var mutex sync.Mutex

func Ping(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	count++
	mutex.Unlock()
	fmt.Fprintf(w, "path is : %v\n", r.URL.Path)
}
func Counter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	fmt.Fprintf(w, "count is : %v\n", count)
	mutex.Unlock()
}
