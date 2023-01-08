//反向代理
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func JsonToMap(fp string) (mapResult map[string]interface{}) {
	jsonStr, _ := ioutil.ReadFile(fp + "config.json")
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}
	//fmt.Println(mapResult)
	return
}

type handles struct {
	uport map[string]interface{}
}

func newhandles(fp string) *handles {
	return &handles{
		uport: JsonToMap(fp),
	}
}
func (h *handles) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Host //不能用r.URL.Host=""
	inp := h.uport[host]
	if inp == nil { //这个不知道为什么，总有nil值访问
		return
	}
	port := inp.(string)      //host对应的 port
	host = "http://127.0.0.1" //host会自动转化为公网ip，需要对port开防火墙。因为是本地直接127.0.0.1，就解决了。
	remote, err := url.Parse(host + ":" + port)
	if err != nil {
		fmt.Println(host+":"+port, err)
		return
		//panic(err)
	}
	/*
		关键的代码就是NewSingleHostReverseProxy这个方法，
		查看源码的话不难看出该方法返回了一个ReverseProxy对象，
		在ReverseProxy中的ServeHTTP方法实现了这个具体的过程，
		主要是对源http包头进行重新封装，而后发送到后端服务器
	*/
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}
