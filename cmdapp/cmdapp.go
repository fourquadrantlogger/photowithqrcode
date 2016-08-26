package main

import (
	"flag"
	"code.aliyun.com/mougew/photowithqrcode/drawqrcode"
	"io/ioutil"
	"os"
)


var (
	bgname=flag.String("bg","bgimg/京东广告图.jpg","二维码的背景图文件路径")
	outpath=flag.String("out","bgimg/京东广告图.png","输出文件路径,请附加后缀png")
	qrtext=flag.String("qrtext","请输入二维码文本","请输入二维码文本")
	x=flag.Int("x",648,"二维码位置x")
	y=flag.Int("y",205,"二维码位置y")
	w=flag.Int("w",220,"二维码宽")
	h=flag.Int("h",220,"二维码高")
)
func main() {
 	flag.Parse()
	bgbs,err:=ioutil.ReadFile(*bgname)
	if err!=nil{
		panic(err)
	}
	img:=drawqrcode.Get(bgbs,*qrtext,*x,*y,*w,*h)

	bf:=drawqrcode.ImageSave(img)
	ioutil.WriteFile(*outpath,bf.Bytes(),os.ModePerm)
}
