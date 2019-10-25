package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			panic(err)
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		file, head, err := r.FormFile("file")
		if err != nil {
			io.WriteString(w, "faild to get data err:%s\n"+err.Error())
			return
		}
		defer file.Close()
		newfile, err := os.Create("./img/" + head.Filename)
		if err != nil {
			fmt.Printf("faild to create file ,err:%s", err.Error())
			return
		}
		defer newfile.Close()
		_, err = io.Copy(newfile, file)
		if err != nil {
			fmt.Printf("faild to save data into file,err:%s", err.Error())
			return
		}
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

//文件上传成功
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "upload file success!")
}
