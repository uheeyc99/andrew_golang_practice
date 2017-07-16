package main

import (
	"net/http"
	"fmt"
)

const SERVER_PORT  = 8080
const SERVER_DOMAIN = "192.168.88.105"
const RESPONSE_TEMPLATE  = "hello"
func rootHandler(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","text/html")
	w.Header().Set("Content-Length",fmt.Sprint(len(RESPONSE_TEMPLATE)))
	fmt.Sprint("1")
	fmt.Sprint(r.Header)
	w.Write([]byte(RESPONSE_TEMPLATE))
}

func main(){

	http.HandleFunc(fmt.Sprintf("%s:%d/",SERVER_DOMAIN,SERVER_PORT),rootHandler)
	err:=http.ListenAndServeTLS(fmt.Sprintf(":%d",SERVER_PORT),"private/ssl.crt","private/ssl.key",nil)
	if err!=nil{
		fmt.Println(err)
	}
}
