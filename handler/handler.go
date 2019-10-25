package handler

import (
	"fmt"
	"github.com/20100204/baiduyun/meta"
	"github.com/20100204/baiduyun/util"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
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
		filemeta := meta.FileMeta{
			FileName: head.Filename,
			Location: "./img/" + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		newfile, err := os.Create(filemeta.Location)
		if err != nil {
			fmt.Printf("faild to create file ,err:%s", err.Error())
			return
		}
		defer newfile.Close()
		filemeta.FileSize, err = io.Copy(newfile, file)
		if err != nil {
			fmt.Printf("faild to save data into file,err:%s", err.Error())
			return
		}
		newfile.Seek(0, 0)
		filemeta.FileSha1 = util.FileSha1(newfile)
		meta.UpdateFileMeta(filemeta)
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

//文件上传成功
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "upload file success!")
}
