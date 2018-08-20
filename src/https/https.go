package main

import(
	"fmt"
	"net/http"
)

const (
	SERVER_PORT = 8888
	SERVER_DOMAIN = "localhost"
	RESPONSE_TEMPLATE = "HELLO"
)

func rootHandler(w http.ResponseWriter, req *http.Request){
	fmt.Println(req)
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", fmt.Sprint(len(RESPONSE_TEMPLATE)))
	w.Write([]byte(RESPONSE_TEMPLATE))
}

func main(){
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", SERVER_PORT), nil)
}