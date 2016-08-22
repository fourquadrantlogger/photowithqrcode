package main

import (
	"io/ioutil"
	"./drawqrcode"
	"net/http"
	"strconv"
	"fmt"
)
var (
	bgname="京东广告图.jpg"
	qrtext="填写测试网址"
	x=640
	y=200
	w=238
	h=238
)
func handler(resp http.ResponseWriter,req *http.Request)  {
	v:=req.URL.Query()
	fmt.Println(req.URL.String(),v)
	bgname:=v["bg"][0]
	qrtext:=v["qrtext"][0]
	x,err:=strconv.Atoi(v["x"][0])
	y,err:=strconv.Atoi(v["y"][0])
	w,err:=strconv.Atoi(v["width"][0])
	h,err:=strconv.Atoi(v["height"][0])

	bgbs,err:=ioutil.ReadFile("bgimg/"+bgname)
	if err!=nil{
		panic(err)
	}
	img:=drawqrcode.Get(bgbs,qrtext,x,y,w,h)
	bf:=drawqrcode.ImageSave(img)

	resp.Write(bf.Bytes())
}

func main() {
	http.HandleFunc("/favicon.ico", func(resp http.ResponseWriter,req *http.Request) {
		resp.WriteHeader(404)
		resp.Write([]byte{})
	})
	http.HandleFunc("/",handler)
	err := http.ListenAndServe(":8005",nil)
	if err!=nil{
		panic(err)
	}

}
