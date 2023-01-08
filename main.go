//反向代理
package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var fp string //程序绝对路径目录，win服务必须绝对路径.
//反向代理
func main() {
	fp = getCurrentAbPath() + "\\"
	h := newhandles(fp)
	err := http.ListenAndServe(":80", h)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

//获取程序绝对路径目录
func getCurrentAbPath() string {
	exePath, err := os.Executable()
	if err != nil {
		//log.Fatal(err)
		return ""
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res

}

/*
func main() {
	fp = getCurrentAbPath() + "\\"
	shs(fp)       //守护程序
	startServer() //反向代理
}


*/
