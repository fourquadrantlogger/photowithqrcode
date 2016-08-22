package main

import (
	"net/http"
	"io/ioutil"
	"strconv"
	"os"
)

func main() {
	for i:=1;i<=34;i++{

		url := "http://localhost:8005/makeimg?bg=抱抱广告图.jpg&x=648&y=207&width=220&height=220&qrtext=http://mxz.so/1u"+strconv.Itoa(i)

		req, _ := http.NewRequest("GET", url, nil)
		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		ioutil.WriteFile("output/baobao_"+strconv.Itoa(i)+".png",body,os.ModePerm)
	}
	for i:=1;i<=34;i++{

		url := "http://localhost:8005/makeimg?bg=京东广告图.jpg&x=648&y=207&width=220&height=220&qrtext=http://mxz.so/2u"+strconv.Itoa(i)

		req, _ := http.NewRequest("GET", url, nil)
		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		ioutil.WriteFile("output/jd_"+strconv.Itoa(i)+".png",body,os.ModePerm)
	}



}