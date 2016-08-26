package main

import (
	"io/ioutil"
	"code.aliyun.com/mougew/photowithqrcode/drawqrcode"
	"net/http"
	"strconv"
	"fmt"
	"log"
	"strings"
)
var (
	bgname="京东广告图.jpg"
	qrtext="填写测试网址"
	x=648
	y=205
	w=220
	h=220
)
func handlerMakeqrimg(resp http.ResponseWriter,req *http.Request)  {
	v:=req.URL.Query()
	fmt.Println(req.URL.String(),v)
	bgname:=v["bg"][0]
	qrtext:=v["qrtext"][0]
	x,err:=strconv.Atoi(v["x"][0])
	y,err:=strconv.Atoi(v["y"][0])
	w,err:=strconv.Atoi(v["width"][0])
	h,err:=strconv.Atoi(v["height"][0])

	bgbs:=make([]byte,0)
	if(strings.Contains(bgname,"http://")||strings.Contains(bgname,"https://")){
		url := bgname
		req, _ := http.NewRequest("GET", url, nil)
		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		bgbs, _ = ioutil.ReadAll(res.Body)
	}else {
		bgbs,err=ioutil.ReadFile("bgimg/"+bgname)
	}

	if(err!=nil){
		resp.Write([]byte(err.Error()))
	}
	img:=drawqrcode.Get(bgbs,qrtext,x,y,w,h)
	bf:=drawqrcode.ImageSave(img)

	resp.Header().Add("Access-Control-Allow-Origin", "*")
	resp.Write(bf.Bytes())
}
func staticDirHandler(mux *http.ServeMux, prefix string) {
	mux.HandleFunc(prefix,
		func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL.Path)
			file :=r.URL.Path[1:]
			log.Println(file)
			http.ServeFile(w, r, file)
		})
}
func main() {
	var mux = http.NewServeMux()
	mux.HandleFunc("/favicon.ico", func(resp http.ResponseWriter,req *http.Request) {
		resp.WriteHeader(404)
		resp.Write([]byte{})
	})
	mux.HandleFunc("/makeimg",handlerMakeqrimg)
	staticDirHandler(mux, "/bgimg/")
	err := http.ListenAndServe(":8005",mux)
	if err!=nil{
		panic(err)
	}

}
