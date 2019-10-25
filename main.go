package main

import (
	"fmt"
	"github.com/20100204/baiduyun/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload",handler.UploadHandler)
	http.HandleFunc("/file/upload/suc",handler.UploadSucHandler)
	err:=http.ListenAndServe(":8089",nil)
	if err != nil {
		fmt.Printf("%s",err)
	}
}
